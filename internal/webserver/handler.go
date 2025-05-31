package webserver

import(
	"fmt"
	"net/http"
	"strconv"

	"github.com/mouboo/cowsayaas/internal/cowsay"
)

// PlainHandler handles the plain text API
func PlainHandler(w http.ResponseWriter, r *http.Request) {
	//Default width
	width := 44
	// If the user provided a "?width=<number>" and it can be parsed
	// to an int, use that number instead.
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
		width = widthParsed
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
