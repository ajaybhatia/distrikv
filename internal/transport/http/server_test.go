package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ajaybhatia/distrikv/internal/storage"
)

// BenchmarkProductionLogger is a placeholder for benchmarking the production logger implementation.
// It currently does not contain any benchmark code and can be implemented in the future to measure
// the performance of the logging system under production-like conditions.
func BenchmarkProductionLogger(b *testing.B) {
	// Initialize our decoupled memory storage engine
	memEngine := storage.NewMemoryEngine()

	// Create a new HTTP server instance, passing the initialized storage engine to it.
	httpServer := NewServer(memEngine)

	// Create a new HTTP request for benchmarking. This is a placeholder request and can be modified
	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/v1/kv/testkey", nil)

	// Reset timer to ignere the initialization time of the server and storage engine.
	b.ResetTimer()

	// Run the benchmark loop, sending the HTTP request to the server and recording the response.
	for i := 0; i < b.N; i++ {
		httpServer.router.ServeHTTP(w, req)
	}
}
