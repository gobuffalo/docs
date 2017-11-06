package main

import (
	"log"

	"github.com/gobuffalo/gobuffalo/actions"
)

func main() {
	log.Fatal(actions.App().Serve())
}
