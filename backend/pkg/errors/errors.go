package errors

import (
	"encoding/json"
	"net/http"
	"nethub-mdm/pkg/logger"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func RespondError(w http.ResponseWriter, log logger.Logger, code int, errType, message string, err error) {
	if err != nil && log != nil {
		log.Errorf("API Error [%s]: %v", errType, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:   errType,
		Message: message,
	})
}
