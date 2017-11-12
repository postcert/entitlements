package main

import (
	"log"

	"github.com/postcert/entitlements/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
