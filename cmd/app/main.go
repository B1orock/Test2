package main

import (
	"Varian_v2/config"
	api "Varian_v2/internal/restapi"
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

	// Инициализация API
	api := api.NewAPI(*cfg)

	// Запуск сервера
	fmt.Println("Запуск сервера на порту", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, api)
}
