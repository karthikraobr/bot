package main

import "github.com/Netflix/go-env"

type Config struct {
	DSN     string `env:"DATABASE_URL,required=true"`
	Address string `env:"Address,default=0.0.0.0:8080"`
}

// Load config from environment
func Load() (Config, error) {
	var cfg Config
	_, err := env.UnmarshalFromEnviron(&cfg)
	return cfg, err
}
