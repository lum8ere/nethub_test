package handler

import (
	"encoding/json"
	"net/http"
	"nethub-mdm/internal/service"
	"nethub-mdm/internal/storage/model"
	"nethub-mdm/pkg/errors"
	"nethub-mdm/pkg/logger"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type LocationHandler struct {
	svc service.LocationService
	log logger.Logger
}

func NewLocationHandler(svc service.LocationService, log logger.Logger) *LocationHandler {
	return &LocationHandler{svc: svc, log: log}
}

func (h *LocationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var loc model.Location
	if err := json.NewDecoder(r.Body).Decode(&loc); err != nil {
		errors.RespondError(w, h.log, http.StatusBadRequest, "invalid_json", "Bad Request", err)
		return
	}
	if err := h.svc.Create(r.Context(), &loc); err != nil {
		errors.RespondError(w, h.log, http.StatusInternalServerError, "db_error", "Failed to create", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(loc)
}

func (h *LocationHandler) List(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit <= 0 {
		limit = 10
	}

	locs, total, err := h.svc.List(r.Context(), name, nil, limit, (page-1)*limit)
	if err != nil {
		errors.RespondError(w, h.log, http.StatusInternalServerError, "db_error", "Failed to list", err)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"data": locs, "total": total})
}

func (h *LocationHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var loc model.Location
	json.NewDecoder(r.Body).Decode(&loc)
	h.svc.Update(r.Context(), id, &loc)
	w.WriteHeader(http.StatusOK)
}

func (h *LocationHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.svc.Delete(r.Context(), id)
	w.WriteHeader(http.StatusNoContent)
}
