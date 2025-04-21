# MCP Server Clipboard

A Model Context Protocol (MCP) server that provides clipboard functionality. This implementation follows the MCP specification and offers tools for reading and writing clipboard content.

## Features

- Simple, lightweight MCP-compliant clipboard server
- Secure API token authentication
- Configurable clipboard size limits
- Tools for clipboard content management:
  - `clipboard_update`: Update clipboard content
  - `clipboard_get`: Retrieve clipboard content
- Configurable port via command-line arguments

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/mcp-server-clipboard.git
cd mcp-server-clipboard

# Install dependencies
go mod download
```

## Usage

### Running the server

```bash
# Run with default port (9001)
go run *.go

# Run with a specific port
go run *.go -port 9000
```

### MCP Tools

The server provides the following MCP tools:

#### clipboard_update

Updates the clipboard content with user input.

- **Input Schema**: `{"type":"object","properties":{"content":{"type":"string"}}}`
- **Example**:
  ```json
  {
    "content": "Hi, MCP!"
  }
  ```

#### clipboard_get

Retrieves the current clipboard content.

- **Input Schema**: `{}`
- **Example Response**:
  ```json
  {
    "type": "text",
    "text": "Hi, MCP!"
  }
  ```

### Authentication

All requests to the MCP server require authentication using the API token:

```
Authorization: Bearer 328db9d4ab39ec9a2eceb2f702f42744
```

## Configuration

The clipboard server can be configured with the following settings:

- **Port**: The port on which the server listens (default: 9001)
- **Clipboard Size**: Maximum size of clipboard content in bytes (default: 10,000 bytes)

## Implementation Details

This server implements the Model Context Protocol (MCP) specification using the `github.com/modelcontextprotocol-ce/go-sdk` library. It uses a synchronous server model with HTTP transport for communication.

## License

This project is licensed under the MIT License - see the LICENSE file for details.