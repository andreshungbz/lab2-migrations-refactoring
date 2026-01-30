// Filename: cmd/api/main.go
// A simple HTTP server with routes and handlers.

// NOTE: Because the instructions specify a function signature of SetupRoutes(mux *http.ServeMux),
// it can only handle route-specific middleware since the in-class middleware signature specifies
// returning http.Handler. The SetupMux function was created to handle global middleware.

package main

import (
	"log"
	"net/http"

	"github.com/andreshungbz/lab2-migrations-refactoring/internal/routes"
)

// Main Program

func main() {
	// Route multiplexer setup
	mux := http.NewServeMux()
	routes.SetupRoutes(mux)            // handler registration & route-specific middleware
	wrappedMux := routes.SetupMux(mux) // global middleware

	// Start local web server on port 4000
	log.Print("CMPS3162 Lab 2 Server Starting on :4000")
	err := http.ListenAndServe(":4000", wrappedMux)
	log.Fatal(err)
}
