package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	BindingAddr string `toml:"binding_addr"`
	DatabaseURL string `toml:"database_url"`
}

func Load(filePath string) (*Config, error) {
	cfg := &Config{}

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(bytes, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
