package main

import (
	"context"
	"fmt"
	"net/http"
	"nethub-mdm/internal/config"
	"nethub-mdm/internal/service"
	"nethub-mdm/internal/storage/query"
	"nethub-mdm/internal/transport/http/handler"
	"nethub-mdm/pkg/db_manager"
	"nethub-mdm/pkg/logger"
	"nethub-mdm/pkg/service_helper"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

type AppHandlers struct {
	Device   *handler.DeviceHandler
	Location *handler.LocationHandler
	Platform *handler.PlatformHandler
	Audit    *handler.AuditHandler
}

// @title NetHub MDM API
// @version 1.0
// @description API сервера управления устройствами
// @host localhost:9000
// @BasePath /
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

			db := mgr.DB()
			q := query.Use(db)

			deviceSvc := service.NewDeviceService(q)

			handlers := &AppHandlers{
				Device:   handler.NewDeviceHandler(deviceSvc, logger),
				Location: handler.NewLocationHandler(service.NewLocationService(q), log),
				Platform: handler.NewPlatformHandler(service.NewPlatformService(q), log),
				Audit:    handler.NewAuditHandler(service.NewAuditService(q), log),
			}

			r, err := initRoutes(handlers)
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
