package main

import (
	"log"
	"net/http"
	"os"

	"github.com/xevimacia/zenbot/internal/handler"
)

func main() {
	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create HTTP server
	mux := http.NewServeMux()

	// Register routes
	handler.RegisterRoutes(mux)

	// Start server
	log.Printf("ZenBot API starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
