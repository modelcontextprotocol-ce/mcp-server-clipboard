# MCP Server Clipboard

A Model Context Protocol (MCP) server that provides clipboard functionality.

## Features

- Implements a clipboard resource for reading and writing clipboard content
- Provides a clipboard settings resource for configuring clipboard behavior
- Port configuration via command-line arguments

## Installation

```bash
go get github.com/modelcontextprotocol-ce/mcp-server-clipboard
```

## Usage

### Running the server

```bash
# Run with default port (8080)
go run main.go

# Run with a specific port
go run main.go -port 9000
```

### API Endpoints

#### Clipboard Resource

- **GET /clipboard** - Get current clipboard content
  - Response: `{"content": "clipboard content"}`

- **PATCH /clipboard** - Update clipboard content
  - Request: `{"content": "new clipboard content"}`
  - Response: `{"content": "new clipboard content"}`

#### Clipboard Settings Resource

- **GET /clipboard_settings** - Get current clipboard settings
  - Response: `{"max_size": 10000}`

- **PATCH /clipboard_settings** - Update clipboard settings
  - Request: `{"max_size": 20000}`
  - Response: `{"max_size": 20000}`

## License

This project is licensed under the MIT License - see the LICENSE file for details.