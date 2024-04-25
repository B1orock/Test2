package config

import (
	"github.com/spf13/viper"
	"os"
)

const (
	yaml = "yaml"
	env  = "env"
)

type Config struct {
	Port    string `json:"port"`
	Message string `json:"message"`
}

func GetYAMLConfig(cfg *Config) error {
	viper.SetConfigFile("config/cfg1.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		return err
	}

	return nil
}

func GetENVConfig(cfg *Config) error {
	cfg.Port = os.Getenv("PORT")
	cfg.Message = os.Getenv("MESSAGE")

	return nil
}

/*func LoadConfig() *Config {
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
*/
