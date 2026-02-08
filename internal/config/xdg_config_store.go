// Package config implements configuration storage
package config

import (
	"encoding/json"
	"log"
	"os"
)

type XDGConfigStore struct {
	Path string
}

func (s XDGConfigStore) Exists() bool {
	var err error
	if _, err = os.Stat(s.Path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
func (s XDGConfigStore) GetManagedPaths() ([]string, error) {
	var err error
	if _, err = os.Stat(s.Path); os.IsNotExist(err) {
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

	data, err := os.ReadFile(s.Path)
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

func (s XDGConfigStore) AddManagedPath(path string) error {
	var err error
	if _, err = os.Stat(s.Path); os.IsNotExist(err) {
		log.Print("No config file exists")
	} else {
		log.Print("Config file exists")
	}
	if os.IsNotExist(err) {
		return err
	}
	if err != nil {
		return err
	}

	data, err := os.ReadFile(s.Path)
	if err != nil {
		return err
	}

	var cfg struct {
		ManagedRepos []string `json:"managedRepos"`
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return err
	}
	cfg.ManagedRepos = append(cfg.ManagedRepos, path)
	newData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(s.Path, newData, 0644); err != nil {
		return err
	}
	return nil
}
