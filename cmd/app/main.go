package main

import (
	"Varian_v2/config"
	"fmt"
	"log"
)

func main() {
	// Выбор типа конфигурации
	cfgType := "env"

	// Загрузка конфигурации
	var cfg config.ENVConfig
	var cfg1 config.YAMLConfig

	switch cfgType {
	case "yaml":
		err := config.GetYAMLConfig(&cfg1)
		if err != nil {
			panic(err)
		}
		log.Println()
	case "env":
		cfg, err := config.GetENVConfig()
		fmt.Println(cfg)
		if err != nil {
			panic(err)
		}

	default:
		log.Fatal("Неверный тип конфигурации")
	}

	// Инициализация бизнес-логики
	//	business.Init(&cfg)

	// Создание API-роутера
	//router := api.CreateRouter(&cfg)

	// Запуск сервера
	fmt.Println("Запуск сервера на порту", cfg.Port, "Сообщение", cfg.Message)
	//	http.ListenAndServe(":"+cfg.Port, router)

	// Запуск сервера через YAMLCongig
	fmt.Println("Запуск сервера на порту", cfg1.Port, "Сообщение", cfg1.Message)
	//	http.ListenAndServe(":"+cfg1.Port, router)

	// Запуск сервера через ENVCongig
	//	http.ListenAndServe(fmt.Sprintf(":%d", cfg2.Port), router)
	//fmt.Println("Server started on port", cfg2.Port)
}
