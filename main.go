package main

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"wt/cmd"
	"wt/internal/domain"
	"wt/internal/git"
	"wt/internal/ui"
)

type XDGConfigStore struct {}

func (s XDGConfigStore) GetManagedPaths() ([]string, error) {
	base, err := os.UserConfigDir()
	if err != nil { return nil, err }
	configPath := path.Join(base, ".wt", "config.json")
	if _, err = os.Stat(configPath); os.IsNotExist(err) {
		log.Print("No config file exists")
	} else {
		log.Print("Config file exists")
	}
	if os.IsNotExist(err) {
		return []string{}, nil
	}
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg struct {
		ManagedRepos []string `json:"managedRepos"`
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return cfg.ManagedRepos, nil
}

func main() {
	deps := domain.Dependencies{
		Git: git.ExecGit{},
		ConfigStore: XDGConfigStore{},
		Renderer: ui.FmtRenderer{},
	}
	cmd := cmd.NewRootCmd(deps)
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
