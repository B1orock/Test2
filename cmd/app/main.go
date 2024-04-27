package main

import (
	"Varian_v2/config"
	"Varian_v2/internal/api"
	"fmt"
	"log"
	"net/http"
)

func main() {

	cfg, err := config.GetENVConfig()
	if err != nil {
		log.Fatal("error load config", err)
	}
	log.Println(cfg)

	// Создание API-роутера
	router := api.CreateRouter(cfg)

	// Запуск сервера
	fmt.Println("Запуск сервера на порту", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, router)
}
