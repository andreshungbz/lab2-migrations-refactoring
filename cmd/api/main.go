// Filename: cmd/api/main.go
// A simple HTTP server with routes and handlers.

package main

import (
	"fmt"
	"log"
	"net/http"
)

// Route struct stores the path and response body text
type Route struct {
	Path string
	Text string
}

// data organizes the route path and text
var data = map[string]Route{
	"home": {
		Path: "/",
		Text: "Hey, Andres here! Welcome to my CMPS3161 Lab 1! My semester project will be on Hotel Room & Housekeeping Management since I recently engaged with the Westin hotel in the Boston Seaport District. It will be fun investigating how hotels in Belize are similar or different.\n",
	},
	"about": {
		Path: "/about",
		Text: "My name is Andres Hung. I'm currently a 2nd semester bachelor's student in IT at the University of Belize. A hobby that I enjoy is drawing. Currently I use a MacBook Air at school and at home my PC dual boots Windows 11 and NixOS. I like computers.\n",
	},
	"contact": {
		Path: "/contact",
		Text: "My student email can be contacted at 2018118240@ub.edu.bz. My GitHub username is andreshungbz. It is the same username that I typically use for other websites too.\n",
	},
	"calculate": {
		Path: "/calculate",
		Text: "The result of 6 + 7 = 13!\n",
	},
}

// Handlers

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(data["home"].Text))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(data["about"].Text))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(data["contact"].Text))
}

// Creative - Calculate uses fmt.Appendf to evaluate a given expression in the returned string.
func calculateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(fmt.Appendf(nil, "The result of 6 + 7 = %d!\n", 6+7))
}

// Main Program

func main() {
	// route multiplexer
	mux := http.NewServeMux()
	// map routes to handlers
	mux.HandleFunc(data["home"].Path, homeHandler)
	mux.HandleFunc(data["about"].Path, aboutHandler)
	mux.HandleFunc(data["contact"].Path, contactHandler)
	mux.HandleFunc(data["calculate"].Path, calculateHandler)

	// start local web server on port 4000
	log.Print("CMPS3162 Lab 1 Server Starting on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
