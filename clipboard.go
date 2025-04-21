package main

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/modelcontextprotocol-ce/go-sdk/spec"
)

// ClipboardHandler implements handlers for clipboard operations
type ClipboardHandler struct {
	mu       sync.RWMutex
	content  string
	settings ClipboardSettings
}

// ClipboardSettings represents the configurable settings for clipboard
type ClipboardSettings struct {
	MaxSize int `json:"max_size"`
}

// NewClipboardHandler creates a new clipboard handler with default settings
func NewClipboardHandler() *ClipboardHandler {
	return &ClipboardHandler{
		settings: ClipboardSettings{
			MaxSize: 10000, // Default 10KB limit
		},
	}
}

// Get returns the current clipboard content
func (c *ClipboardHandler) Get(ctx context.Context, uri string) ([]byte, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return (&spec.TextContent{Text: c.content}).MarshalJSON()
}

// Patch handles updating the clipboard content
func (c *ClipboardHandler) Patch(ctx context.Context, params []byte) (interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Convert data to string
	content := string(params)
	if content == "" {
		return nil, errors.New("content must be a string")
	}

	// Validate content size
	if len(content) > c.settings.MaxSize {
		return nil, fmt.Errorf("content exceeds maximum size of %d bytes", c.settings.MaxSize)
	}

	// Update clipboard content
	c.content = content

	return &spec.TextContent{
		Text: c.content,
	}, nil
}
