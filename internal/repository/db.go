package repository

import (
	"github.com/kelseyhightower/envconfig"
)

type Connect struct {
	Port    string `envconfig:"PORT"`
	Message string `envconfig:"MESSAGE"`
	DB      struct {
		Host     string `envconfig:"DB_HOST" default:"localhost"`
		Port     int    `envconfig:"DB_PORT" default:"5432"`
		User     string `envconfig:"DB_USER" default:"Kostua"`
		Password string `envconfig:"DB_PASSWORD" default:"pgp2fasm3p"`
		Name     string `envconfig:"DB_NAME" default:"SQL-intro"`
	}
}

func GetENVConfig() (*Connect, error) {
	var cfg Connect
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
