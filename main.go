package main

import (
	"crypto/tls"
	"encoding/binary"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"fserver_scanner/pb"

	"google.golang.org/protobuf/proto"
)

type ServerInfo struct {
	Resources []string `json:"resources"`
	// Add other fields as needed
}

type ScannerConfig struct {
	Locale         string
	MinPlayers     int
	MaxPlayers     int
	ResourceSearch string
	EndpointFilter string
	ShowResources  bool
	ShowVars       bool
	ShowFullStruct bool
}

func main() {
	// Define command-line flags
	config := ScannerConfig{}

	flag.StringVar(&config.Locale, "locale", "", "Filter by locale (e.g., 'de-DE')")
	flag.IntVar(&config.MinPlayers, "min-players", 0, "Minimum number of players")
	flag.IntVar(&config.MaxPlayers, "max-players", 0, "Maximum number of players (0 = no limit)")
	flag.StringVar(&config.ResourceSearch, "resource", "", "Search for servers with specific resource (partial match)")
	flag.StringVar(&config.EndpointFilter, "endpoint", "", "Filter by endpoint containing this string")
	flag.BoolVar(&config.ShowResources, "show-resources", true, "Show resource list")
	flag.BoolVar(&config.ShowVars, "show-vars", false, "Show server variables")
	flag.BoolVar(&config.ShowFullStruct, "show-full", false, "Show full server struct")

	flag.Parse()

	url := "https://servers-frontend.fivem.net/api/servers/streamRedir/"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to fetch: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Reading servers from protobuf stream...")
	printFilters(config)
	fmt.Println()

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	serverCount := 0
	// Read loop (depending on how messages are framed, usually length-delimited)
	for {
		// Read the message size prefix (varint or fixed size depends on framing)
		// This is crucial because protobuf messages are not self-delimited.
		// Let's assume each message is prefixed by a 4-byte length header (big endian):

		var lengthBuf [4]byte
		_, err := io.ReadFull(resp.Body, lengthBuf[:])
		if err != nil {
			if err == io.EOF {
				break // End of stream
			}
			fmt.Fprintf(os.Stderr, "Failed to read message length: %v\n", err)
			break
		}

		length := int(binary.LittleEndian.Uint32(lengthBuf[:]))
		if length <= 0 {
			fmt.Fprintf(os.Stderr, "Invalid message length: %d\n", length)
			continue
		}

		msgBuf := make([]byte, length)
		_, err = io.ReadFull(resp.Body, msgBuf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read message body: %v\n", err)
			break
		}

		var server pb.Server
		err = proto.Unmarshal(msgBuf, &server)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to unmarshal protobuf: %v\n", err)
			continue
		}

		// Apply filters
		if server.Data != nil && matchesFilters(server, config) {
			serverCount++
			printServer(server, config, client)
		}
	}

	fmt.Printf("\nFound %d servers matching criteria.\n", serverCount)
}

func printFilters(config ScannerConfig) {
	fmt.Println("Filters:")
	if config.Locale != "" {
		fmt.Printf("  - Locale: %s\n", config.Locale)
	}
	if config.MinPlayers > 0 {
		fmt.Printf("  - Min Players: %d\n", config.MinPlayers)
	}
	if config.MaxPlayers > 0 {
		fmt.Printf("  - Max Players: %d\n", config.MaxPlayers)
	}
	if config.ResourceSearch != "" {
		fmt.Printf("  - Resource Search: %s\n", config.ResourceSearch)
	}
	if config.EndpointFilter != "" {
		fmt.Printf("  - Endpoint Filter: %s\n", config.EndpointFilter)
	}
}

func matchesFilters(server pb.Server, config ScannerConfig) bool {
	data := server.Data
	if data == nil {
		return false
	}

	// Check locale filter
	if config.Locale != "" {
		locale := data.GetVars()["locale"]
		if locale != config.Locale {
			return false
		}
	}

	// Check player count filters
	clients := int(data.GetClients())
	if config.MinPlayers > 0 && clients < config.MinPlayers {
		return false
	}
	if config.MaxPlayers > 0 && clients > config.MaxPlayers {
		return false
	}

	// Check endpoint filter
	if config.EndpointFilter != "" {
		connectEndpoints := data.GetConnectEndPoints()
		hasEndpoint := false
		for _, endpoint := range connectEndpoints {
			if strings.Contains(endpoint, config.EndpointFilter) {
				hasEndpoint = true
				break
			}
		}
		if !hasEndpoint {
			return false
		}
	}

	// Check resource search (requires fetching info.json)
	if config.ResourceSearch != "" {
		if !hasResource(server, config.ResourceSearch) {
			return false
		}
	}

	return true
}

func hasResource(server pb.Server, resourceSearch string) bool {
	connectEndpoints := server.Data.GetConnectEndPoints()
	if len(connectEndpoints) == 0 {
		return false
	}

	endpoint := connectEndpoints[0]
	endpoint = strings.TrimSuffix(endpoint, "/")

	var infoURL string
	if len(endpoint) > 7 && (endpoint[:7] == "http://" || endpoint[:8] == "https://") {
		infoURL = fmt.Sprintf("%s/info.json", endpoint)
	} else {
		infoURL = fmt.Sprintf("http://%s/info.json", endpoint)
	}

	client := &http.Client{
		Timeout: 3 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Get(infoURL)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	var serverInfo ServerInfo
	if err := json.NewDecoder(resp.Body).Decode(&serverInfo); err != nil {
		return false
	}

	// Check if any resource contains the search string
	for _, resource := range serverInfo.Resources {
		if strings.Contains(strings.ToLower(resource), strings.ToLower(resourceSearch)) {
			return true
		}
	}

	return false
}

func printServer(server pb.Server, config ScannerConfig, client *http.Client) {
	data := server.Data

	fmt.Printf("=== Server: %s (%d/%d) ===\n",
		data.GetHostname(), data.GetClients(), data.GetSvMaxclients())
	fmt.Printf("Endpoint: %s\n", server.GetEndPoint())
	fmt.Printf("Gametype: %s\n", data.GetGametype())
	fmt.Printf("Map: %s\n", data.GetMapname())

	if config.ShowFullStruct {
		fmt.Printf("Full Server Struct:\n%+v\n", server)
	}

	// Fetch and show resources if requested
	if config.ShowResources {
		connectEndpoints := data.GetConnectEndPoints()
		if len(connectEndpoints) > 0 {
			endpoint := connectEndpoints[0]
			endpoint = strings.TrimSuffix(endpoint, "/")
			var infoURL string
			if len(endpoint) > 7 && (endpoint[:7] == "http://" || endpoint[:8] == "https://") {
				infoURL = fmt.Sprintf("%s/info.json", endpoint)
			} else {
				infoURL = fmt.Sprintf("http://%s/info.json", endpoint)
			}

			resp, err := client.Get(infoURL)
			if err != nil {
				fmt.Printf("Failed to fetch resources: %v\n", err)
			} else {
				defer resp.Body.Close()

				if resp.StatusCode == http.StatusOK {
					var serverInfo ServerInfo
					if err := json.NewDecoder(resp.Body).Decode(&serverInfo); err != nil {
						fmt.Printf("Failed to parse resources: %v\n", err)
					} else {
						fmt.Printf("Resources (%d):\n", len(serverInfo.Resources))
						/*for _, res := range serverInfo.Resources {
							fmt.Printf("  - %s\n", res)
						}*/
					}
				}
			}
		}
	}

	if config.ShowVars {
		fmt.Printf("Vars:\n")
		for key, value := range data.GetVars() {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}

	fmt.Println()
}
