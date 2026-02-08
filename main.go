package main

import (
	"os"
	"path"
	"wt/cmd"
	"wt/internal/config"
	"wt/internal/domain"
	"wt/internal/git"
	"wt/internal/ui"
)

func userConfigPath() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	configPath := path.Join(base, ".wt", "config.json")
	if _, err = os.Stat(configPath); os.IsNotExist(err) {
		return "", err
	} else {
		return configPath, nil
	}
}

func main() {
	path, _ := userConfigPath()
	deps := domain.Dependencies{
		Git: git.ExecGit{},
		ConfigStore: config.JSONConfigStore{
			Path: path,
		},
		Renderer: ui.FmtRenderer{},
	}
	cmd := cmd.NewRootCmd(deps)
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
