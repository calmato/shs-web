package main

import (
	"os"

	cmd "github.com/calmato/shs-web/api/internal/classroom/cmd/server"
)

func main() {
	if err := cmd.Exec(); err != nil {
		os.Exit(1)
	}
}
