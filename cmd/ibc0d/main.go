package main

import (
	"os"

	"github.com/test/ibc0/cmd/ibc0d/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
