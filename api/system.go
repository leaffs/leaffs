package api

import (
	"leaffs/utils"
	"net/http"
)

func NewSystemHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterSystemRoutes(route *http.ServeMux) {
	route.HandleFunc("GET /register", h.SysInfo)
}

func (h *Handler) SysInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	utils.WriteJson(w, map[string]string{
		"success": "true",
		"message": "System API is working",
	})
}
