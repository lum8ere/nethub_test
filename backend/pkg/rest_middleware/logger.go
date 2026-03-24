package rest_middleware

import (
	"net/http"
	"nethub-mdm/pkg/logger"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func StructuredLogger(log logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()

			defer func() {
				log.Infof("HTTP Request: method=%s path=%s status=%d duration=%s ip=%s",
					r.Method, r.URL.Path, ww.Status(), time.Since(t1), r.RemoteAddr)
			}()

			next.ServeHTTP(ww, r)
		})
	}
}
