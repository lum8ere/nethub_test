package handler

import (
	"encoding/json"
	"net/http"
	"nethub-mdm/internal/service"
	"nethub-mdm/pkg/logger"
	"strconv"
)

type AuditHandler struct {
	svc service.AuditService
	log logger.Logger
}

func NewAuditHandler(svc service.AuditService, log logger.Logger) *AuditHandler {
	return &AuditHandler{svc: svc, log: log}
}

func (h *AuditHandler) List(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit <= 0 {
		limit = 20
	}

	logs, total, err := h.svc.List(r.Context(), limit, (page-1)*limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":  logs,
		"total": total,
	})
}
