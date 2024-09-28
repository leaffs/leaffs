package api

import (
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Register API is working"})
	}
}
