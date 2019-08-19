package main

import (
	"log"

	"github.com/gobuffalo/docs/actions"
)

func main() {
	app := actions.App()
	log.Fatal(app.Serve())
}
