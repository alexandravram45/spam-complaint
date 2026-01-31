package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Spam struct {
		Keywords []string `yaml:"keywords"`
	} `yaml:"spam"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
