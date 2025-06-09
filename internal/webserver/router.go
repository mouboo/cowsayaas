package webserver

import (
	"net/http"
	
	//"github.com/mouboo/cowsayaas/internal/assets"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	
	// Serve cowsays with a plain text http interface
	mux.HandleFunc("/plain", PlainHandler)
	
	// Serve docs from static html files
	docsFileServer := http.FileServer(http.Dir("./assets/docs"))
	mux.Handle("/docs/", http.StripPrefix("/docs", docsFileServer))
	
	// Redirect /docs to /docs/ so index.html is served
	mux.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/", http.StatusMovedPermanently)
	})
	
	//mux.HandleFunc("/docs", DocsHandler)
	
	return mux
}
