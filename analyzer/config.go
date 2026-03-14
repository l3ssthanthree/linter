package analyzer

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ExtraSensitivePatterns []string `yaml:"extra_sensitive_patterns"`
}

func loadConfig(path string) (Config, error) {
	var cfg Config

	if path == "" {
		return cfg, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return cfg, nil
		}
		return cfg, err
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
