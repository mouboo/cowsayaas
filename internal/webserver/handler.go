package webserver

import(
	"fmt"
	"net/http"
	"strconv"

	"github.com/mouboo/cowsayaas/internal/cowsay"
)

// TODO: fix bug if width is less than 5 (messes up slice boundaries)
// Fix could be to let width refer to text width sans borders?
func PlainHandler(w http.ResponseWriter, r *http.Request) {
	//Default width
	width := 44
	// If the user provided a "?width=<number>" and it can be parsed
	// to an int, use that number instead.
	widthStr := r.URL.Query().Get("width")
	if widthStr != "" {
		if widthParsed, err := strconv.Atoi(widthStr); err == nil {
			if widthParsed > 0 {
				width = widthParsed
			} else {
				http.Error(w, "Width must be a positive number", http.StatusBadRequest)
				return
			}
		} else {
			http.Error(w, "Invalid width parameter", http.StatusBadRequest)
			return
		}
	}
	
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "Missing text parameter", http.StatusBadRequest)
		return
	}
	response := cowsay.RenderCowsay(text, width)
	// Write to the ResponseWriter
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, response)
}
