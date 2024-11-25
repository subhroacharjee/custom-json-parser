package main

import (
	"log"

	clirunner "github.com/subhroacharjee/custom-json-parser/internal/cli_runner"
)

func main() {
	if err := clirunner.Run(); err != nil {
		log.Panic(err)
	}
}
