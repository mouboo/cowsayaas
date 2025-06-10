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

	// Figure out what kind of request it is, and populate the cowspec c
	// with the use of helper functions.
	if r.Method == http.MethodGet {
		err = parseFromQuery(r, &c)
	} else if r.Method == http.MethodPost {
		contentType := r.Header.Get("Content-Type")
		if strings.HasPrefix(contentType, "application/json") {
			err = parseFromJSON(r, &c)
		} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
			err = parseFromForm(r, &c)
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

func parseFromQuery(r *http.Request, c *cowspec.CowSpec) error {
	var err error

	// Parse text
	if v := r.URL.Query().Get("text"); v != "" {
		c.Text = v
	}
	// Parse width
	if v := r.URL.Query().Get("width"); v != "" {
		width, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		c.Width = width
	}
	// Parse file
	if v := r.URL.Query().Get("file"); v != "" {
		c.File = v
	}
	// Parse eyes
	if v := r.URL.Query().Get("eyes"); v != "" {
		c.Eyes = v
	}
	// Parse tongue
	if v := r.URL.Query().Get("tongue"); v != "" {
		c.Tongue = v
	}
	return err
}

// Parsing request data from JSON into the cowspec
func parseFromJSON(r *http.Request, c *cowspec.CowSpec) error {
	// Make sure the reader closes before the function ends
	defer r.Body.Close()
	// Read the (JSON) body into the variable body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	// Fill in the struct at c with the data from the JSON request body
	err = json.Unmarshal(body, c)
	// If all went well, err is nil
	return err
}

// Parsing request data from form urlencoded into the cowspec
func parseFromForm(r *http.Request, c *cowspec.CowSpec) error {
	var err error

	if err := r.ParseForm(); err != nil {
		return err
	}

	if v := r.Form.Get("text"); v != "" {
		c.Text = v
	}
	if v := r.Form.Get("width"); v != "" {
		if width, err := strconv.Atoi(v); err == nil {
			c.Width = width
		}
	}
	if v := r.Form.Get("think"); v != "" {
		if v == "true" || v == "True" {
			c.Think = true
		}
	}
	if v := r.Form.Get("file"); v != "" {
		c.File = v
	}
	if v := r.Form.Get("mode"); v != "" {
		c.Mode = v
	}
	if v := r.Form.Get("eyes"); v != "" {
		c.Eyes = v
	}
	if v := r.Form.Get("tongue"); v != "" {
		c.Tongue = v
	}
	return err
}
