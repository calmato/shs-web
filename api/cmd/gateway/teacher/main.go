package main

import (
	"log"
	"os"

	cmd "github.com/calmato/shs-web/api/internal/gateway/cmd/teacher"
)

func main() {
	if err := cmd.Exec(); err != nil {
		log.Printf("An error has occurred: %v", err)
		os.Exit(1)
	}
}
