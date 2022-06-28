package config

import (
	"blacklist-service/database"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host     string            `yaml:"host"`
	Port     int               `yaml:"port"`
	Database *database.Options `yaml:"database"`
}

func Init(path string) (*Config, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	file, err := ioutil.ReadFile(abs)
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}

	if cfg.Host == "" {
		cfg.Host = "0.0.0.0"
	}

	if cfg.Port == 0 {
		cfg.Port = 8080
	}

	return &cfg, nil
}
