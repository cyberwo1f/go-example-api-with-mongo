package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port int `env:"PORT,required"`
	DB   *databaseConfig
}

type databaseConfig struct {
	Database string `env:"MONGODB_DATABASE,required"`
	URL      string `env:"MONGODB_URL,required"`
	Source   string `env:"MONGODB_SOURCE"`
}

func LoadConfig(ctx context.Context) (*Config, error) {
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) Address() string {
	return fmt.Sprintf(":%d", cfg.Port)
}
