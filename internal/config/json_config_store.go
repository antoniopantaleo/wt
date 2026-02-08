// Package config implements configuration storage
package config

import (
	"encoding/json"
	"log"
	"os"
)

type JSONConfigStore struct {
	Path string
}

type storeConfig struct {
	ManagedRepos []string `json:"managedRepos"`
}

func (s JSONConfigStore) Exists() bool {
	_, err := os.Stat(s.Path)
	return err == nil
}

func (s JSONConfigStore) GetManagedPaths() ([]string, error) {
	exists, err := s.checkConfigFile()
	if os.IsNotExist(err) || !exists {
		return []string{}, nil
	}
	if err != nil {
		return nil, err
	}

	cfg, err := s.readConfig()
	if err != nil {
		return nil, err
	}
	return cfg.ManagedRepos, nil
}

func (s JSONConfigStore) AddManagedPath(path string) error {
	exists, err := s.checkConfigFile()
	if !exists {
		return err
	}
	if err != nil {
		return err
	}

	cfg, err := s.readConfig()
	if err != nil {
		return err
	}

	cfg.ManagedRepos = append(cfg.ManagedRepos, path)
	return s.writeConfig(cfg)
}

func (s JSONConfigStore) checkConfigFile() (bool, error) {
	_, err := os.Stat(s.Path)
	if os.IsNotExist(err) {
		log.Print("No config file exists")
		return false, err
	}
	if err != nil {
		return false, err
	}

	log.Print("Config file exists")
	return true, nil
}

func (s JSONConfigStore) readConfig() (storeConfig, error) {
	data, err := os.ReadFile(s.Path)
	if err != nil {
		return storeConfig{}, err
	}

	var cfg storeConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return storeConfig{}, err
	}

	return cfg, nil
}

func (s JSONConfigStore) writeConfig(cfg storeConfig) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.Path, data, 0644)
}
