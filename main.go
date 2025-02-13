package main

import (
	"fmt"
	"log"
	"os"

	"core-lnurl/config"
	"core-lnurl/server"
)

func main() {
	// Default config path
	configPath := "config.yml"

	// Allow overriding config file from command-line argument
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	// Load config
	if err := config.LoadConfig(configPath); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Print config (for debugging)
	fmt.Printf("Loaded Config: %+v\n", config.AppConfig)

	// Start the server
	server.StartServer()
}
