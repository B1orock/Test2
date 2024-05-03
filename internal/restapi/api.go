package api

import (
	"Varian_v2/config"
	"Varian_v2/internal/routes"
	"net/http"
)

// Структура API
type API struct {
	cfg config.Config
}

// Создание нового API
func NewAPI(cfg config.Config) *API {
	return &API{cfg: cfg}
}

// Обработка GET запроса /get_message
func (api *API) GetMessage(w http.ResponseWriter, r *http.Request) {
	routes.GetMessage(w, r, api.cfg.Message)
}

// Сервер API
func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		api.GetMessage(w, r)
	default:
		http.NotFound(w, r)
	}
}
