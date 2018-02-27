package main

import (
	"log"

	"github.com/gobuffalo/gobuffalo/actions"
)

func main() {
	app := actions.App()
	log.Fatal(app.Serve())
}
