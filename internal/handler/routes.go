package handler

import "net/http"

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
	// Register the /zenbot endpoint with POST method restriction
	mux.HandleFunc("/zenbot", MethodMiddleware(http.MethodPost)(HandleZenbotRequest))
}
