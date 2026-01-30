// Filename: internal/middleware/middleware.go
// Contains various middleware functions for the routes.

package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the HTTP method, path, and timestamp for a request.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Pre-Processing
		log.Printf("[LOG] %v %v", r.Method, r.URL.Path) // default log behavior provides the timestamp

		// Call-Handler
		next.ServeHTTP(w, r)
	})
}

// TimingMiddleware calculates and logs the elapsed time for the next handler called.
func TimingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Pre-Processing
		start := time.Now() // record current time before handler call

		// Call-Handler
		next.ServeHTTP(w, r)

		// Post-Processing
		elapsed := time.Since(start) // record elapsed time since start
		log.Printf("[TIME] Route (%v %v) processed in %v", r.Method, r.URL.Path, elapsed)
	})
}
