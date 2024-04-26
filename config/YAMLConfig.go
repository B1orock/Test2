package config

import "github.com/spf13/viper"

type YAMLConfig struct {
	Port    string `yaml:"port"`
	Message string `yaml:"message"`
}

func GetYAMLConfig(cfg1 *YAMLConfig) error {
	viper.SetConfigFile("config/cfg1.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(cfg1)
	if err != nil {
		return err
	}

	return nil
}
