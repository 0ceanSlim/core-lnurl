package lightning

import (
	"fmt"
	"log"
	"net/url"

	"core-lnurl/config"

	"github.com/gorilla/websocket"
)

// CLNClient manages the WebSocket connection to a Core Lightning node
type CLNClient struct {
	conn *websocket.Conn
}

// NewCLNClient initializes a WebSocket connection to CLN using a Rune
func NewCLNClient() (*CLNClient, error) {
	clnHost := config.AppConfig.Lightning.Host
	rune := config.AppConfig.Lightning.Rune

	// Build WebSocket URL with Rune authentication
	wsURL := url.URL{
		Scheme: "ws",
		Host:   clnHost,
		Path:   "/v1/commando",
	}

	headers := make(map[string][]string)
	headers["X-Commando-Auth"] = []string{rune}

	// Connect to WebSocket
	conn, _, err := websocket.DefaultDialer.Dial(wsURL.String(), headers)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to CLN WebSocket: %w", err)
	}

	return &CLNClient{conn: conn}, nil
}

// SendCommand sends a command to the CLN node
func (c *CLNClient) SendCommand(method string, params map[string]interface{}) (map[string]interface{}, error) {
	// Build JSON-RPC request
	request := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  method,
		"params":  params,
	}

	// Send request
	if err := c.conn.WriteJSON(request); err != nil {
		return nil, fmt.Errorf("failed to send command: %w", err)
	}

	// Read response
	var response map[string]interface{}
	if err := c.conn.ReadJSON(&response); err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check for errors
	if errorData, exists := response["error"]; exists {
		return nil, fmt.Errorf("CLN error: %v", errorData)
	}

	return response, nil
}

// Close closes the WebSocket connection
func (c *CLNClient) Close() {
	if err := c.conn.Close(); err != nil {
		log.Printf("Error closing WebSocket: %v", err)
	}
}
