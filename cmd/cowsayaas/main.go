package main

import (
	"fmt"

	"github.com/mouboo/cowsayaas/internal/cowsay"
)

func main() {
	s := cowsay.RenderCowsay("Moo!")
	fmt.Println(s)
}
