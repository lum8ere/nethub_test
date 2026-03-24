package main

import (
	"context"
	"fmt"
	"net/http"
	"nethub-mdm/internal/config"
	"nethub-mdm/internal/transport/http/handler"
	"nethub-mdm/pkg/db_manager"
	"nethub-mdm/pkg/logger"
	"nethub-mdm/pkg/service_helper"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

type AppHandlers struct {
	Device *handler.DeviceHandler
}

func main() {
	log, err := logger.NewZapLogger("api")
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer log.Sync()

	if os.Getenv("ENV_PATH") == "" {
		os.Setenv("ENV_PATH", "../../../.env")
	}

	// 3. Загружаем конфигурацию
	cfg := config.LoadConfig(log)

	var allRoutes *chi.Mux
	var webServer *http.Server

	err = service_helper.StartService("api",
		log,
		cfg.DatabaseURL,
		func(ctx context.Context, logger logger.Logger, mgr db_manager.Manager) error {
			r, err := initRoutes()
			if err != nil {
				return err
			}
			allRoutes = r
			return nil
		},
		func(ctx context.Context, logger logger.Logger, mgr db_manager.Manager) error {
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
		func(ctx context.Context, logger logger.Logger, mgr db_manager.Manager) error {
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

	if err != nil {
		log.Fatalf("Service stopped with fatal error: %v", err)
	}
}
