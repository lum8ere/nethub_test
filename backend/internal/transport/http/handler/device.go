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

type DeviceHandler struct {
	svc service.DeviceService
	log logger.Logger
}

func NewDeviceHandler(svc service.DeviceService, log logger.Logger) *DeviceHandler {
	return &DeviceHandler{svc: svc, log: log}
}

// @Summary Создать устройство
// @Tags devices
// @Accept json
// @Produce json
// @Param request body model.Device true "Данные устройства"
// @Success 201 {object} model.Device
// @Failure 400 {object} errors.ErrorResponse
// @Router /api/v1/devices [post]
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

// @Summary Список устройств
// @Tags devices
// @Produce json
// @Param hostname query string false "Поиск по hostname"
// @Param is_active query boolean false "Фильтр по активности"
// @Param page query int false "Номер страницы" default(1)
// @Param limit query int false "Количество на странице" default(10)
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/devices [get]
func (h *DeviceHandler) List(w http.ResponseWriter, r *http.Request) {
	hostname := r.URL.Query().Get("hostname")
	isActiveStr := r.URL.Query().Get("is_active")

	var isActive *bool
	if isActiveStr != "" {
		val := isActiveStr == "true"
		isActive = &val
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 || limit > 100 {
		limit = 10
	}
	offset := (page - 1) * limit

	devices, count, err := h.svc.List(r.Context(), hostname, isActive, limit, offset)
	if err != nil {
		errors.RespondError(w, h.log, http.StatusInternalServerError, "db_error", "Ошибка при получении списка", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":  devices,
		"total": count,
		"page":  page,
		"limit": limit,
	})
}

// @Summary Получить устройство по ID
// @Tags devices
// @Produce json
// @Param id path string true "ID устройства (UUID)"
// @Success 200 {object} model.Device
// @Failure 400 {object} errors.ErrorResponse "Неверный ID"
// @Failure 404 {object} errors.ErrorResponse "Устройство не найдено"
// @Router /api/v1/devices/{id} [get]
func (h *DeviceHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	dev, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "device not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(dev)
}

// @Summary Обновить устройство
// @Tags devices
// @Accept json
// @Produce json
// @Param id path string true "ID устройства (UUID)"
// @Param request body model.Device true "Обновленные данные"
// @Success 200 {object} model.Device
// @Failure 400 {object} errors.ErrorResponse "Неверный запрос"
// @Failure 500 {object} errors.ErrorResponse "Ошибка сервера"
// @Router /api/v1/devices/{id} [put]
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

// @Summary Удалить устройство (Soft Delete)
// @Tags devices
// @Produce json
// @Param id path string true "ID устройства (UUID)"
// @Success 204 "Успешно удалено (No Content)"
// @Failure 400 {object} errors.ErrorResponse "Неверный ID"
// @Failure 500 {object} errors.ErrorResponse "Ошибка сервера"
// @Router /api/v1/devices/{id} [delete]
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
