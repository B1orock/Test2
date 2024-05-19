package routes

import (
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

// Обработка GET запроса /get_message
func GetMessage(w http.ResponseWriter, r *http.Request, message string) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", message)
}
