// Filename: cmd/api/main.go
// A simple HTTP server with routes, handlers, and middleware.

// NOTE: Because the instructions specify a function signature of SetupRoutes(mux *http.ServeMux),
// it can only handle route-specific middleware since the in-class middleware signature specifies
// returning http.Handler. The SetupMux function was created to handle global middleware.

package main

import (
	"log"
	"net/http"

	"github.com/andreshungbz/lab2-migrations-refactoring/internal/routes"
)

func main() {
	// Route multiplexer setup
	mux := http.NewServeMux()
	routes.SetupRoutes(mux)            // handler registration & route-specific middleware
	wrappedMux := routes.SetupMux(mux) // global middleware

	// Start local web server on port 4000
	log.Print("[CMPS3162-LAB-02] HTTP Server Starting @ http://localhost:4000/")
	err := http.ListenAndServe(":4000", wrappedMux)
	log.Fatal(err)
}
