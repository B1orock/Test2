package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port     string `envconfig:"PORT"`
	DBConfig DBConfig
	//	Mes      Mes
	Message string `envconfig:"MESSAGE"`
}
type DBConfig struct {
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     int    `envconfig:"DB_PORT" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

/*
type Mes struct {
	Message string `envconfig:"MESSAGE"`
}
*/

func GetENVConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
