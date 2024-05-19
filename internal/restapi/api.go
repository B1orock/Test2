package api

import (
	"Varian_v2/config"
	"Varian_v2/internal/repository"
	"Varian_v2/internal/routes"
	"encoding/json"
	"net/http"
)

// Структура API
type API struct {
	cfg  config.Config
	repo *repository.Repository
	//msg  config.Mes
}

// Создание нового API
func NewAPI(cfg config.Config, repo *repository.Repository) *API {
	return &API{cfg: cfg, repo: repo}
}

// Обработка GET запроса /get_message
func (api *API) GetMessage(w http.ResponseWriter, r *http.Request) {
	routes.GetMessage(w, r, api.cfg.Message)
}

// Обработка GET запроса /contacts
func (api *API) GetContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := api.repo.GetContact()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(contacts)
}

// Сервер API
func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		switch r.URL.Path {
		case "/get_message":
			api.GetMessage(w, r)
		case "/get_all_data":
			api.GetContacts(w, r)
		default:
		}
	}
}
