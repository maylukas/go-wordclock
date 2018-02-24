package http_api

import (
	"net/http"
	"encoding/json"
	"github.com/maylukas/go-wordclock/backend/repository/config_repository"
)

func updateConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		c := config_repository.WordclockConfig{}
		json.NewDecoder(r.Body).Decode(&c)
		config_repository.SetConfig(c)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
