package cowsay

import (
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	
	// Serve cowsays with a plain text http interface
	mux.HandleFunc("/api", APIHandler)
	
	// Serve docs from static html files
	docsFileServer := http.FileServer(http.Dir("./assets/docs"))
	mux.Handle("/docs/", http.StripPrefix("/docs", docsFileServer))
	
	// Redirect /docs to /docs/ so index.html is served
	mux.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/", http.StatusMovedPermanently)
	})
	
	
	return mux
}
