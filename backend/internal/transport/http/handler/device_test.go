package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"nethub-mdm/internal/storage/model"
	"nethub-mdm/pkg/logger"
	"nethub-mdm/pkg/types"
)

type MockDeviceService struct {
	CreateFunc func(ctx context.Context, dev *model.Device) error
}

func (m *MockDeviceService) Create(ctx context.Context, dev *model.Device) error {
	return m.CreateFunc(ctx, dev)
}
func (m *MockDeviceService) List(ctx context.Context, h string, ia *bool, l, o int) ([]*model.Device, int64, error) {
	return nil, 0, nil
}
func (m *MockDeviceService) GetByID(ctx context.Context, id string) (*model.Device, error) {
	return nil, nil
}
func (m *MockDeviceService) Update(ctx context.Context, id string, dev *model.Device) error {
	return nil
}
func (m *MockDeviceService) Delete(ctx context.Context, id string) error {
	return nil
}

func TestDeviceHandler_Create(t *testing.T) {
	log, _ := logger.NewZapLogger("test")

	t.Run("success create", func(t *testing.T) {
		mockSvc := &MockDeviceService{
			CreateFunc: func(ctx context.Context, dev *model.Device) error {
				dev.ID = types.StrPtr("uuid-123")
				return nil
			},
		}

		h := NewDeviceHandler(mockSvc, log)

		body := map[string]string{
			"hostname":      "test-host",
			"ip":            "10.0.0.1",
			"platform_code": "LINUX",
		}
		jsonBody, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/devices", bytes.NewBuffer(jsonBody))
		w := httptest.NewRecorder()

		h.Create(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("expected status 201, got %d", w.Code)
		}

		var resp model.Device
		json.NewDecoder(w.Body).Decode(&resp)
		if *resp.ID != "uuid-123" {
			t.Errorf("expected ID uuid-123, got %s", *resp.ID)
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		h := NewDeviceHandler(&MockDeviceService{}, log)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/devices", bytes.NewBufferString("invalid"))
		w := httptest.NewRecorder()

		h.Create(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", w.Code)
		}
	})
}
