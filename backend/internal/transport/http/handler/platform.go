package handler

import (
	"encoding/json"
	"net/http"
	"nethub-mdm/internal/service"
	"nethub-mdm/pkg/logger"
)

type PlatformHandler struct {
	svc service.PlatformService
	log logger.Logger
}

func NewPlatformHandler(svc service.PlatformService, log logger.Logger) *PlatformHandler {
	return &PlatformHandler{svc: svc, log: log}
}

func (h *PlatformHandler) List(w http.ResponseWriter, r *http.Request) {
	platforms, err := h.svc.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(platforms)
}
