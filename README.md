# FiveM Server Scanner

A Go application that connects to the FiveM server stream API, decodes protobuf messages, and displays real-time server information with advanced filtering options.

## Description

This application fetches server data from the FiveM server stream API endpoint, decodes the protobuf responses, and displays server information in real-time. It supports advanced filtering and inspection of server resources, endpoints, and variables.

## Features

- Connects to FiveM's server stream API (protobuf-based)
- Decodes and parses protobuf server data
- Real-time server information display
- Advanced filtering by locale, player count, resource, and endpoint
- Optionally fetches and displays server resources and variables
- Error handling for network, protobuf, and parsing issues

## Requirements

- Go 1.21 or higher

## Usage

1. Clone or download this repository
2. Navigate to the project directory
3. Run the application with optional filters:

```bash
go run main.go [flags]
```

### Command-line Flags

- `-locale string`         Filter by locale (e.g., 'de-DE')
- `-min-players int`       Minimum number of players
- `-max-players int`       Maximum number of players (0 = no limit)
- `-resource string`       Search for servers with a specific resource (partial match)
- `-endpoint string`       Filter by endpoint containing this string
- `-show-resources`        Show resource list (default: true)
- `-show-vars`             Show server variables
- `-show-full`             Show full server struct

## How it works

The application:
1. Connects to the FiveM server stream API and reads a stream of length-prefixed protobuf messages
2. Decodes each message into a `Server` structure (see `pb/server.proto`)
3. Applies user-specified filters (locale, player count, resource, endpoint)
4. Optionally fetches additional server info (resources, variables) from each server's `/info.json`
5. Displays the filtered server information to the console

## Output

The application will continuously stream and display server information matching the filters until the connection is closed or an error occurs. Each server's data is printed to the console, with optional details based on flags.

## Error Handling

The application includes error handling for:
- Network connection issues
- HTTP/protobuf response errors
- Protobuf and JSON parsing errors
- Scanner and filter errors 