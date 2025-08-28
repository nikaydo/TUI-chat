package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Env struct {
	ThemePath string `env:"THEME_PATH"`
	LangPath  string `env:"LANG_PATH"`

	TCP addr `envPrefix:"TCP_"`
	UDP addr `envPrefix:"UDP_"`
}

type addr struct {
	Port int    `env:"PORT"`
	Host string `env:"HOST"`
}

func ReadEnv() (Env, error) {
	err := godotenv.Load()
	var cfg Env
	if err != nil {
		return cfg, fmt.Errorf("error read env: %w", err)
	}
	if err := env.Parse(&cfg); err != nil {
		return cfg, fmt.Errorf("error read env: %w", err)
	}
	return cfg, nil
}
