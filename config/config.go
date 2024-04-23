package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port string `json:"port"`
}

func LoadConfig() *Config {
	// Считать конфигурацию из файла
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	// Распарсить JSON
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
