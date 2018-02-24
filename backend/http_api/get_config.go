package http_api

import (
	"net/http"
	"encoding/json"
	"github.com/maylukas/go-wordclock/backend/repository/config_repository"
)

func getCurrentConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c := config_repository.GetConfig()
		json.NewEncoder(w).Encode(c)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
