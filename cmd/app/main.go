package main

import (
	"fmt"
	"net/http"

	"Varian_v2/api"
	"Varian_v2/busines"
	"Varian_v2/config"
)

func main() {
	// Загрузить конфигурацию
	cfg := config.LoadConfig()

	// Инициализировать бизнес-логику
	busines.Init(cfg)

	// Создать API-роутер
	router := api.CreateRouter(cfg)

	// Запустить сервер
	fmt.Println("Запуск сервера на порту", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, router)
}
