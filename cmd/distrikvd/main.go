package main

import (
	"log"

	"github.com/ajaybhatia/distrikv/internal/storage"
	"github.com/ajaybhatia/distrikv/internal/transport/http"
)

// main is the entry point of the DistriKV server application. It initializes the core engine modules and starts the server.
func main() {
	log.Println("Initializing DistriKV Core Engine Modules...")

	// Initialize the in-memory storage engine. This will be used to store key-value pairs in memory.
	memEngine := storage.NewMemoryEngine()

	// Create a new HTTP server instance, passing the initialized storage engine to it.
	httpServer := http.NewServer(memEngine)

	log.Println("Starting DistriKV HTTP Server on :8080...")
	// Start the HTTP server on port 8080. If there is an error starting the server,
	// log the error and exit the application.
	if err := httpServer.Start(":8080"); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
