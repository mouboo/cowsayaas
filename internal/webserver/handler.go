package webserver

import(
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mouboo/cowsayaas/internal/cowsay"
	"github.com/mouboo/cowsayaas/internal/spec"
)

// PlainHandler handles the plain text API
func PlainHandler(w http.ResponseWriter, r *http.Request) {
	// First make sure the request is valid (GET)
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Fill in a CowSpec with all the options
	c := spec.NewCowSpec()
	
	// Text is a required parameter
	c.Text = r.URL.Query().Get("text")
	if c.Text == "" {
		http.Error(w, "Missing text parameter", http.StatusBadRequest)
		return
	}
	
	// Width is an optional parameter, representing the maximum width
	// of the text (sans borders) displayed. Default: 40.
	widthStr := r.URL.Query().Get("width")	
	if widthStr != "" {
		widthParsed, err := strconv.Atoi(widthStr)
		if err != nil {
			http.Error(w, "Invalid width parameter", http.StatusBadRequest)
			return
		}
		if widthParsed < 1 {
			http.Error(w, "width must be a positive number", http.StatusBadRequest)
			return
		}
		c.Width = widthParsed
	}
	
	response, err := cowsay.RenderCowsay(c)
	if err != nil {
		log.Printf("RenderCowsay error: %v", err)
		http.Error(w, "Internal server error in rendering", http.StatusInternalServerError)
		return
	}
	// Write to the ResponseWriter
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, response)
	return
}
