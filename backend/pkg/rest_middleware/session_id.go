package rest_middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const sessionIDKey contextKey = "session_id"

func GenerateSessionID() string {
	return uuid.New().String()
}

func SessionID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID := r.Header.Get("X-Session-Id")
		if sessionID == "" {
			sessionID = GenerateSessionID()
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, sessionIDKey, sessionID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetSessionID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if sessionID, ok := ctx.Value(sessionIDKey).(string); ok {
		return sessionID
	}
	return ""
}
