package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gobuffalo/gobuffalo/actions"
)

func main() {
	fmt.Println("PORT", os.Getenv("PORT"))
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
