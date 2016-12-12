package main

import (
	"log"
	"net/http"

	"github.com/markbates/gobuffalo/actions"
)

func main() {
	log.Fatal(http.ListenAndServe(":3000", actions.App()))
}
