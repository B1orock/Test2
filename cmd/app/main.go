package main

import (
	"Varian_v2/config"
	"Varian_v2/internal/repository"
	api "Varian_v2/internal/restapi"
	"context"
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

	// Подключение к БД
	repo, err := repository.NewRepository(context.Background(), cfg.DBConfig)
	if err != nil {
		panic(err)
	}

	// Проверка соединения с БД
	err = repo.Ping(context.Background())
	if err != nil {
		fmt.Println("Ошибка при проверке соединения:", err)
		return
	}
	fmt.Println("Соединение с базой данных установлено успешно!")

	// Инициализация API
	api := api.NewAPI(*cfg, repo)

	// Запуск сервера
	fmt.Println("Запуск сервера на порту", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, api)
}
