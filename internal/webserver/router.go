package webserver

import (
	"net/http"
	
	//"github.com/mouboo/cowsayaas/internal/assets"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	
	// Serve cowsays with a plain text http interface
	mux.HandleFunc("/plain", PlainHandler)
	
	docsFileServer := http.FileServer(http.Dir("./assets/docs"))
	mux.Handle("/docs/", http.StripPrefix("/docs", docsFileServer))
	
	// Serve a help page through embedded static html
	//mux.Handle("/docs/", http.StripPrefix("/docs", http.FileServer(http.FS(assets.DocsFS()))
	
	// Redirect /docs to /docs/ so index.html is served
	mux.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/", http.StatusMovedPermanently)
	})
	
	//mux.HandleFunc("/docs", DocsHandler)
	
	return mux
}
