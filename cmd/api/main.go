// Filename: cmd/api/main.go
// - A simple HTTP server with routes, handlers, and middleware.
package main

// NOTE: Because the instructions specify a function signature of SetupRoutes(mux *http.ServeMux),
// it can only handle route-specific middleware since the in-class middleware signature specifies
// returning http.Handler. The SetupMux function was created to handle global middleware.

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andreshungbz/lab2-migrations-refactoring/internal/routes"
)

func main() {
	// FLAGS

	// Server port
	port := flag.Int("port", 4000, "Port to start the local server on.")
	flag.Parse()
	if !(*port >= 1 && *port <= 65535) {
		fmt.Fprintf(os.Stderr, "[ERROR] Port %d invalid. Port must be 1 - 65535.\n", *port)
	}

	// MAIN

	// Route multiplexer setup
	mux := http.NewServeMux()
	routes.SetupRoutes(mux)            // handler registration & route-specific middleware
	wrappedMux := routes.SetupMux(mux) // global middleware

	// Start local server
	log.Printf("[CMPS3162-LAB-02] HTTP Server Starting @ http://localhost:%d/", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), wrappedMux)
	log.Fatal(err)
}
