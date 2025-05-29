package webserver

import(
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/cowsay", CowsayHandler)
	return mux
}
