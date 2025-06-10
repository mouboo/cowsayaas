package webserver

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/mouboo/cowsayaas/internal/cowsay"
	"github.com/mouboo/cowsayaas/internal/cowspec"
)

// ApiHandler handles requests in various forms, urlencoded, JSON, etc.
func ApiHandler(w http.ResponseWriter, r *http.Request) {
	c := cowspec.NewCowSpec()
	var err error

	// Figure out what kind of request it is, and populate the cowspec
	// with the use of helper functions.
	if r.Method == http.MethodGet {
		c, err = parseFromQuery(r)
	} else if r.Method == http.MethodPost {
		contentType := r.Header.Get("Content-Type")
		if strings.HasPrefix(contentType, "application/json") {
			err = parseFromJSON(r, &c)
		} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
			c, err = parseFromForm(r)
		} else {
			http.Error(w, "Unsupported Content-Type", http.StatusUnsupportedMediaType)
			return
		}
	} else {
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
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

func parseFromQuery(r *http.Request) (cowspec.CowSpec, error) {
	c := cowspec.NewCowSpec()

	// Text is an optional string parameter, defaults to "Moo!"
	c.Text = r.URL.Query().Get("text")
	if c.Text == "" {
		c.Text = "Moo!"
	}

	// Width is an optional integer parameter, representing the maximum width
	// of the text (sans borders) displayed. Default: 40.
	widthStr := r.URL.Query().Get("width")
	if widthStr != "" {
		widthParsed, err := strconv.Atoi(widthStr)
		if err != nil {
			//http.Error(w, "Invalid width parameter", http.StatusBadRequest)
			return c, err
		}
		if widthParsed < 1 {
			//http.Error(w, "width must be a positive number", http.StatusBadRequest)
			return c, err
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

	return c, nil
}

func parseFromJSON(r *http.Request, c *cowspec.CowSpec) error {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, c)
	return err
}

func parseFromForm(r *http.Request) (cowspec.CowSpec, error) {
	c := cowspec.NewCowSpec()
	return c, nil
}
