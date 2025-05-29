package webserver

import(
	"fmt"
	"net/http"

	"github.com/mouboo/cowsayaas/internal/cowsay"
)

// TODO: connect with cowsay renderer
func CowsayHandler(w http.ResponseWriter, r *http.Request) {
	// Render "ascii" art
	s := "lorem ipsum dolor sit amet, consectitur dolor sit amet"
	response := cowsay.RenderCowsay(s)
	// Write to the ResponseWriter
	fmt.Fprint(w, response)
}
