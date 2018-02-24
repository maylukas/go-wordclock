package http_api

import "net/http"

func Configure() {
	http.HandleFunc("/api/get_config", getCurrentConfig)
	http.HandleFunc("/api/set_config", updateConfig)
	http.ListenAndServe(":80", nil)
}
