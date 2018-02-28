package http_api

import (
	"net/http"
	"fmt"
)

func Configure() {
	srv := &http.Server{Addr:":8080"}
	http.HandleFunc("/api/get_config", getCurrentConfig)
	http.HandleFunc("/api/set_config", updateConfig)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Printf("Starting http server failed: %v", err)
	}
}
