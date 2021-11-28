package main

import (
	"log"

	"github.com/iforgot2021/golang/tour/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute error: %v\n", err)
	}
}
