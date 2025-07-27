package handler

import "net/http"

// RegisterRoutes registers all HTTP routes for the ZenBot API
func RegisterRoutes(mux *http.ServeMux) {
	// Register the /zenbot endpoint
	mux.HandleFunc("POST /zenbot", HandleZenbotRequest)
}
