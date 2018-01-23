package main

import (
	"log"

	"github.com/gobuffalo/gobuffalo/actions"
	"github.com/kr/pretty"
)

func main() {
	a := actions.App()
	pretty.Println("### a.Options ->", a.Options)
	log.Fatal(a.Serve())
}
