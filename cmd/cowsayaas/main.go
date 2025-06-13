// Main package for Cowsay as a Service
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	
	"github.com/mouboo/cowsayaas/internal/cowsay"
)

var version = "0.1"

func main() {
	// Determine port to listen on. Defaults to 8080.
	port := "8080"
	
	// Environment variable COWPORT overrides default.
	if envPort := os.Getenv("COWPORT"); envPort != "" {
		port = envPort
	}
	
	// Flag -port overrides both default and environment.
	flagPort := flag.String("port", port, "Port to listen on")
	flag.Parse()
	port = *flagPort

	// Set up multiplexer for HTTP requests, routing them
	// to different handlers
	mux := cowsay.SetupRoutes()

	// Start the web server
	log.Printf("Starting server version %v on port %v", version, port)
	log.Fatal(http.ListenAndServe(":" + port, mux))
}
