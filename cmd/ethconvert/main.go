package main

import (
	"log"

	"github.com/jon4hz/ethconvert/cmd/ethconvert/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
