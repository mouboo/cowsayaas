package webserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mouboo/cowsayaas/internal/cowsay"
	"github.com/mouboo/cowsayaas/internal/cowspec"
)

// DocsHandler handles serving embedded static html files
func DocsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "test DocsHandler")
}

// PlainHandler handles the plain text API
func ApiHandler(w http.ResponseWriter, r *http.Request) {
	// First make sure the request is valid (GET)
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Fill in a CowSpec with all the options
	c := spec.NewCowSpec()

	// Text is a required string parameter
	c.Text = r.URL.Query().Get("text")
	if c.Text == "" {
		http.Error(w, "Missing text parameter", http.StatusBadRequest)
		return
	}

	// Width is an optional integer parameter, representing the maximum width
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

	// File is optional, defaults to "default"
	tmpFile := r.URL.Query().Get("file")
	if tmpFile != "" {
		c.File = tmpFile
	}

	// Modes set both eyes and tongue. Can be individually overridden with eyes
	// and/or tongue parameters.
	// borg, dead, greedy, paranoia, stoned, tired, wired, youthful
	switch r.URL.Query().Get("mode") {
	case "borg":
		c.Eyes = "=="
	case "dead":
		c.Eyes = "xx"
		c.Tongue = "U"
	case "greedy":
		c.Eyes = "$$"
	case "paranoia":
		c.Eyes = "@@"
	case "stoned":
		c.Eyes = "**"
		c.Tongue = "U"
	case "tired":
		c.Eyes = "--"
	case "wired":
		c.Eyes = "OO"
	case "youthful":
		c.Eyes = ".."
	}

	// Eyes is an optional parameter, if not set the template cow-file will
	// fill in a default
	tmpEyes := r.URL.Query().Get("eyes")
	if tmpEyes != "" {
		c.Eyes = tmpEyes
	}

	// Tongue is an optional parameter, if not set the template cow-file will
	// fill in a default
	tmpTongue := r.URL.Query().Get("tongue")
	if tmpTongue != "" {
		c.Tongue = tmpTongue
	}

	// Render the cowsay according to the cowspec
	response, err := cowsay.RenderCowsay(c)
	if err != nil {
		log.Printf("RenderCowsay error: %v", err)
		http.Error(w, "Internal server error in rendering", http.StatusInternalServerError)
		return
	}
	// Write to the ResponseWriter
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, response)
	log.Println("response sent")
	return
}
