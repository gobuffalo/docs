package main

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/gobuffalo/actions"
)

func main() {
	port := envy.Get("PORT", "3000")
	log.Fatal(actions.App().Start(port))
}
