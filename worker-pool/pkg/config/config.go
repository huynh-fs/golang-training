package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	WorkerPool struct {
		Workers int `yaml:"workers"`
		Tasks   int `yaml:"tasks"`
	} `yaml:"worker_pool"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}