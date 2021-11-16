package main

import (
	"os"

	cmd "github.com/calmato/shs-web/api/internal/gateway/cmd/teacher"
)

func main() {
	if err := cmd.Exec(); err != nil {
		os.Exit(1)
	}
}
