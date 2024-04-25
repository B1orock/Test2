package main

import (
	"Varian_v2/api"
	"Varian_v2/businnes"
	"Varian_v2/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Выбор типа конфигурации
	cfgType := "env"

	// Загрузка конфигурации
	var cfg config.Config
	switch cfgType {
	case "yaml":
		err := config.GetYAMLConfig(&cfg)
		if err != nil {
			panic(err)
		}
		log.Println()
	case "env":
		err := config.GetENVConfig(&cfg)
		if err != nil {
			panic(err)
		}
	default:
		log.Fatal("Неверный тип конфигурации")
	}

	// Инициализация бизнес-логики
	business.Init(&cfg)

	// Создание API-роутера
	router := api.CreateRouter(&cfg)

	// Запуск сервера
	fmt.Println("Запуск сервера на порту", cfg.Port, "Сообщение", cfg.Message)
	http.ListenAndServe(":"+cfg.Port, router)
}
