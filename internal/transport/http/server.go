package http

import (
	"errors"
	"net/http"

	"github.com/ajaybhatia/distrikv/internal/storage"
	"github.com/gin-gonic/gin"
)

// Server represents the HTTP server for Distrikv,
// handling incoming requests and interacting with the storage engine.
type Server struct {
	engine storage.StorageEngine
	router *gin.Engine
}

// NewServer initializes and returns a new Server instance with the provided storage engine.
func NewServer(engine storage.StorageEngine) *Server {
	// Set Gin to release mode for production use, which disables debug output and optimizes performance.
	gin.SetMode(gin.ReleaseMode)
	// Create a new Gin router instance to handle HTTP requests.
	r := gin.New()
	r.Use(gin.Recovery()) // Use Gin's recovery middleware to handle panics and prevent server crashes.

	// Initialize the Server struct with the provided storage engine and the Gin router.
	s := &Server{
		engine: engine,
		router: r,
	}

	// Set up the HTTP routes for the server.
	s.setupRoutes()

	// Return the initialized Server instance.
	return s
}

// setupRoutes defines the HTTP routes and their corresponding handler functions for the server.
func (s *Server) setupRoutes() {
	s.router.POST("/v1/kv/:key", s.handlePut) // Route for handling PUT requests to store key-value pairs.
	s.router.GET("/v1/kv/:key", s.handleGet)
}

// handlePut processes incoming PUT requests to store key-value pairs in the storage engine.
func (s *Server) handlePut(c *gin.Context) {
	key := c.Param("key")        // Extract the key from the URL parameter.
	value, err := c.GetRawData() // Read the raw request body to get the value to be stored.

	// If there is an error reading the request body, respond with a 400 Bad Request status and return.
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Attempt to store the key-value pair in the storage engine. If there is an error during storage,
	if err := s.engine.Put(c.Request.Context(), key, value); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // If there is an error storing the key-value pair, respond with a 500 Internal Server Error status and return the error message in JSON format.
		return
	}

	c.Status(http.StatusCreated) // If the key-value pair is stored successfully, respond with a 201 Created status.
}

// handleGet processes incoming GET requests to retrieve values associated with keys from the storage engine.
func (s *Server) handleGet(c *gin.Context) {
	key := c.Param("key") // Extract the key from the URL parameter.

	// Attempt to retrieve the value associated with the key from the storage engine. If there is an error during retrieval,
	value, err := s.engine.Get(c.Request.Context(), key)
	if err != nil {
		// If the error is due to the key not being found in the storage engine, respond with a 404 Not Found status. Otherwise, respond with a 500 Internal Server Error status and return the error message in JSON format.
		if errors.Is(err, storage.ErrKeyNotFound) {
			c.Status(http.StatusNotFound)
			return
		}
		// For any other error during retrieval, respond with a 500 Internal Server Error status and return the error message in JSON format.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", value) // If the value is retrieved successfully, respond with a 200 OK status and return the value as binary data.
}

// Start runs the HTTP server on the specified address.
func (s *Server) Start(addr string) error {
	return s.router.Run(addr) // Start the Gin router on the specified address and return any error that occurs during startup.
}
