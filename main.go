package main

import (
	"log"

	"github.com/mehrdadjalili/facegram_auth_service/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
