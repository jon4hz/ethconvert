package main

import (
	"log"

	"github.com/jon4hz/ethconvert/cmd/ethconvert/cmd"
)

func main() {
	/* 	amount := decimal.NewFromInt(50)
	   	gwei, err := ethconvert.Convert(amount, "ether", "gwei")
	   	if err != nil {
	   		panic(err)
	   	}
	   	fmt.Println(gwei) */
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
