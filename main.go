package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/gobuffalo/actions"
)

func main() {
	port := envy.Get("PORT", "3000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), actions.App()))
}
