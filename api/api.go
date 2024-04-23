package api

import (
	"fmt"
	"net/http"

	"Varian_v2/business"
	"Varian_v2/config"
)

func CreateRouter(cfg *config.Config) http.Handler {
	router := http.NewServeMux()

	// Определить маршрут для приветствия
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := business.GetGreeting()
		fmt.Fprintf(w, message)
	})

	return router
}
