package webserver

import(
	"fmt"
	"net/http"
	"strconv"

	"github.com/mouboo/cowsayaas/internal/cowsay"
)

func PlainHandler(w http.ResponseWriter, r *http.Request) {
	//Default width
	width := 44
	// If the user provided a "?width=<number>" and it can be parsed
	// to an int, use that number instead.
	widthStr := r.URL.Query().Get("width")
	if widthStr != "" {
		if w, err := strconv.Atoi(widthStr); err == nil {
			width = w
		}
	}
	
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "Missing text parameter", http.StatusBadRequest)
	}
	response := cowsay.RenderCowsay(text, width)
	// Write to the ResponseWriter
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, response)
}
