package handler

import (
	"encoding/json"
	"net/http"
	"nethub-mdm/internal/service"
	"nethub-mdm/internal/storage/model"

	"github.com/go-chi/chi/v5"
)

type DeviceHandler struct {
	svc service.DeviceService
}

func NewDeviceHandler(svc service.DeviceService) *DeviceHandler {
	return &DeviceHandler{svc: svc}
}

func (h *DeviceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dev model.Device
	if err := json.NewDecoder(r.Body).Decode(&dev); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.svc.Create(r.Context(), &dev); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dev)
}

func (h *DeviceHandler) List(w http.ResponseWriter, r *http.Request) {
	hostname := r.URL.Query().Get("hostname")
	isActiveStr := r.URL.Query().Get("is_active")

	var isActive *bool
	if isActiveStr != "" {
		val := isActiveStr == "true"
		isActive = &val
	}

	devices, err := h.svc.List(r.Context(), hostname, isActive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(devices)
}

func (h *DeviceHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	dev, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "device not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(dev)
}

func (h *DeviceHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "invalid device id", http.StatusBadRequest)
		return
	}

	var dev model.Device
	if err := json.NewDecoder(r.Body).Decode(&dev); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.svc.Update(r.Context(), idStr, &dev); err != nil {
		http.Error(w, "failed to update device: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dev)
}

func (h *DeviceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "invalid device id", http.StatusBadRequest)
		return
	}

	if err := h.svc.Delete(r.Context(), idStr); err != nil {
		http.Error(w, "failed to delete device: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
