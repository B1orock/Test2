package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port    string `envconfig:"PORT"`
	Message string `envconfig:"MESSAGE"`
}

func GetENVConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	//fmt.Println("Port: ", cfg.Port)
	//fmt.Printf("Message: %s\n", cfg.Message)
	return &cfg, nil
}
