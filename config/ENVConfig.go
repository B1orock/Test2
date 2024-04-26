package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type ENVConfig struct {
	Port    int    `envconfig:"PORT"`
	Message string `envconfig:"MESSAGE"`
}

func GetENVConfig() (*ENVConfig, error) {
	var cfg ENVConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	fmt.Println(cfg.Port)
	fmt.Printf("Message: %s\n", cfg.Message)
	return &cfg, nil
}
