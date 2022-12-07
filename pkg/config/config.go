package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port int `env:"PORT,required"`
	DB   *databaseConfig
	Cors *corsConfig
}

type databaseConfig struct {
	Database string `env:"MONGODB_DATABASE,required"`
	URL      string `env:"MONGODB_URL,required"`
	Source   string `env:"MONGODB_SOURCE"`
}

type corsConfig struct {
	AllowedOrigins     string `env:"ALLOWED_ORIGINS,default=*"`
	AllowedHeaders     string `env:"ALLOWED_HEADERS,default=*"`
	AllowedMethods     string `env:"ALLOWED_METHODS,default=POST, GET, OPTIONS, DELETE, PUT"`
	AllowedCredentials string `env:"ALLOWED_CREDENTIALS,default=true"`
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
