package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database struct {
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
		Name     string `env:"DB_NAME"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
	} `yaml:"database"`

	Token struct {
		AccessSecret string `env:"ACCESS_SECRET"`
	} `yaml:"token"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
