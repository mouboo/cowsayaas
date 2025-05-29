package webserver

import(
	"fmt"
	"net/http"
)

// TODO: connect with cowsay renderer
func CowsayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<cowmessage placeholder>")
}
