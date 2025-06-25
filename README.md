# FiveM Server Scanner

A Go application that connects to the FiveM server stream API and displays real-time server information.

## Description

This application fetches server data from the FiveM server stream API endpoint and parses the JSON responses to display server information in real-time.

## Features

- Connects to FiveM's server stream API
- Parses JSON server data
- Real-time server information display
- Error handling for network and parsing issues

## Requirements

- Go 1.21 or higher

## Usage

1. Clone or download this repository
2. Navigate to the project directory
3. Run the application:

```bash
go run main.go
```

## How it works

The application:
1. Makes an HTTP request to the FiveM server stream API
2. Reads the response as a stream of JSON objects
3. Parses each JSON line into a `ServerInfo` structure
4. Displays the server information to the console

## Output

The application will continuously stream server information until the connection is closed or an error occurs. Each server's data is printed to the console.

## Error Handling

The application includes error handling for:
- Network connection issues
- HTTP response errors
- JSON parsing errors
- Scanner errors 