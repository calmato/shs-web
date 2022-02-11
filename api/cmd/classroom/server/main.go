package main

import (
	"log"
	"os"

	cmd "github.com/calmato/shs-web/api/internal/classroom/cmd/server"
)

func main() {
	if err := cmd.Exec(); err != nil {
		log.Printf("An error has occurred: %v", err)
		os.Exit(1)
	}
}
