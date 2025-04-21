package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/modelcontextprotocol-ce/go-sdk/server"
	"github.com/modelcontextprotocol-ce/go-sdk/server/stream"
	"github.com/modelcontextprotocol-ce/go-sdk/spec"
)

func main() {
	// Parse command line arguments for port configuration
	port := flag.Int("port", 9001, "Port to run the MCP server on")
	flag.Parse()

	// Create clipboard handler
	clipboard := NewClipboardHandler()

	addr := fmt.Sprintf(":%d", *port)

	provider := stream.NewHTTPServerTransportProvider(addr)

	// Create a new MCP server
	srv := server.NewSync(provider).
		WithCapabilities(spec.NewServerCapabilitiesBuilder().Resources(true, false, false).Tools(true, false).Build()).
		WithServerInfo(spec.Implementation{Name: "MCP Clipboard Server", Version: "1.0.0"}).
		WithTools(
			spec.Tool{
				Name: "clipboard_update", Description: "Update clipboard with user input",
				InputSchema: []byte(`{"type":"object","properties":{"content":{"type":"string"}}}`),
			},
			spec.Tool{
				Name: "clipboard_get", Description: "Get current clipboard content",
				InputSchema: []byte(`{}`),
			},
		).
		WithAPIToken("328db9d4ab39ec9a2eceb2f702f42744").(server.SyncBuilder).
		WithToolHandler("clipboard_update", clipboard.Patch).
		WithToolHandler("clipboard_get", clipboard.Get).
		Build()

	// Start the server
	go func() {
		log.Printf("Starting MCP server on: %s", addr)
		if err := srv.Start(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for termination signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down MCP server...")
	srv.Stop()
	log.Println("Server stopped")
}
