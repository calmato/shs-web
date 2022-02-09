package main

import (
	"fmt"
	"os"

	cmd "github.com/calmato/shs-web/api/internal/messenger/cmd/server"
)

func main() {
	if err := cmd.Exec(); err != nil {
		fmt.Printf("An error has occurred: %v", err)
		os.Exit(1)
	}
}
