package main

import (
	"log"

	"github.com/manicar2093/join-matic/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
