package main

import (
	"context"
	"net/http"
	"nethub-mdm/pkg/logger"
	"nethub-mdm/pkg/service_helper"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	var allRoutes *chi.Mux
	var webServer *http.Server

	service_helper.StartService("api",
		func(ctx context.Context, logger logger.Logger) error {
			r, err := initRoutes()
			if err != nil {
				return err
			}
			allRoutes = r
			return nil
		},
		func(ctx context.Context, logger logger.Logger) error {
			webServer = &http.Server{
				Addr:    ":9000",
				Handler: allRoutes,
			}
			logger.Info("Server listening on port 9000")

			go func() {
				if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Error("Server error", "error", err)
				}
			}()

			return nil
		},
		func(ctx context.Context, logger logger.Logger) error {
			shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer shutdownCancel()

			if err := webServer.Shutdown(shutdownCtx); err != nil {
				logger.Error("Error shutting down server", "error", err)
				return err
			}

			logger.Info("Exiting process: Server shut down")
			return nil
		},
	)
}
