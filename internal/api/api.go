package api

import (
	"Varian_v2/config"
	"Varian_v2/internal/service"
	"fmt"
	"net/http"
)

func CreateRouter(cfg *config.Config) http.Handler {
	router := http.NewServeMux()

	// Определить маршрут для приветствия
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := service.GetGreeting()
		fmt.Fprintf(w, message)
	})

	return router
}
