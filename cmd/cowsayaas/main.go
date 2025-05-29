// Main package for Cowsay as a Service
package main

import (
	"log"
	"net/http"

	"github.com/mouboo/cowsayaas/internal/webserver"
)

func main() {

	// Set up multiplexer for HTTP requests, routing them
	// to different handlers
	mux := webserver.SetupRoutes()

	// Start the web server
	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
