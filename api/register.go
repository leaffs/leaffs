package api

import (
	"leaffs/utils"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(route *http.ServeMux) {
	route.HandleFunc("POST /register", h.Register)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	utils.WriteJson(w, map[string]string{
		"success": "true",
		"message": "Register API is working",
	})
}
