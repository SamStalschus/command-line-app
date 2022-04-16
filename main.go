package main

import (
	"command-line/app"
	"log"
	"os"
)

func main() {
	app := app.Generate()
	if error := app.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
