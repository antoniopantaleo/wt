package main

import (
	"os"

	"wt/cmd"
	"wt/internal/domain"
	"wt/internal/git"
	"wt/internal/ui"
)


func main() {
	deps := domain.Dependencies{
		Git: git.ExecGit{},
		Renderer: ui.FmtRenderer{},
	}
	cmd := cmd.NewRootCmd(deps)
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
