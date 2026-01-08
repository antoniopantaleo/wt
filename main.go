package main

import (
	"os"

	"wt/cmd"
	"wt/internal/app"
)

func main() {
	deps := app.Dependencies{}
	cmd := cmd.NewRootCmd(deps)
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
