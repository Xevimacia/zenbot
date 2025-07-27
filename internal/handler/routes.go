package handler

import "net/http"

// SSEMiddleware sets the required headers for Server-Sent Events
func SSEMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Cache-Control")
		next(w, r)
	}
}

// MethodMiddleware creates a middleware that restricts requests to a specific HTTP method
func MethodMiddleware(method string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			next(w, r)
		}
	}
}

// RegisterRoutes registers all HTTP routes for the ZenBot API
func RegisterRoutes(mux *http.ServeMux) {
	// Register the /zenbot endpoint with POST method restriction and SSE headers
	mux.HandleFunc("/zenbot", MethodMiddleware(http.MethodPost)(SSEMiddleware(HandleZenbotRequest)))
}
