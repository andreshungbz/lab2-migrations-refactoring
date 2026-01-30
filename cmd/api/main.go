// Filename: cmd/api/main.go
// A simple HTTP server with routes and handlers.

package main

import (
	"log"
	"net/http"

	"github.com/andreshungbz/lab2-migrations-refactoring/internal/handlers"
	"github.com/andreshungbz/lab2-migrations-refactoring/internal/routes"
)

// Main Program

func main() {
	// route multiplexer
	mux := http.NewServeMux()
	// map routes to handlers
	mux.HandleFunc(routes.Data["home"].Path, handlers.HomeHandler)
	mux.HandleFunc(routes.Data["about"].Path, handlers.AboutHandler)
	mux.HandleFunc(routes.Data["contact"].Path, handlers.ContactHandler)
	mux.HandleFunc(routes.Data["calculate"].Path, handlers.CalculateHandler)

	// start local web server on port 4000
	log.Print("CMPS3162 Lab 2 Server Starting on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
