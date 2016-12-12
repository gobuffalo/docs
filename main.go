package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/markbates/gobuffalo/actions"
	"github.com/markbates/going/defaults"
)

func main() {
	port := defaults.String(os.Getenv("PORT"), "3000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), actions.App()))
}
