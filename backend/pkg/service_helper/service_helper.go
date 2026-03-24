package service_helper

import (
	"context"
	"fmt"
	"nethub-mdm/internal/storage"
	"nethub-mdm/pkg/db_manager"
	"nethub-mdm/pkg/logger"
	"nethub-mdm/pkg/shutdown"
)

type LifecycleFunc func(ctx context.Context, log logger.Logger, mgr db_manager.Manager) error

func StartService(
	serviceName string,
	logger logger.Logger,
	dsn string,
	initFunc LifecycleFunc,
	startFunc LifecycleFunc,
	shutdownFunc LifecycleFunc,
) error {
	prefix := fmt.Sprintf("Service '%s'", serviceName)
	logger.Infof("%s: Initializing", prefix)
	defer logger.Infof("%s: Exited", prefix)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mgr, err := db_manager.NewDbManager(dsn)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	err = mgr.UsePlugins(
		storage.NewAuditPlugin(logger),
	)
	if err != nil {
		return fmt.Errorf("initialization plugins failed: %w", err)
	}

	logger.Info("Starting initialization...")

	if err := initFunc(ctx, logger, mgr); err != nil {
		return fmt.Errorf("Initialization failed: %w", err)
	}

	if err := startFunc(ctx, logger, mgr); err != nil {
		return fmt.Errorf("Start failed: %w", err)
	}

	logger.Info("Service is running. Waiting for shutdown signal...")
	sig := shutdown.WaitForSignalToShutdown()
	cancel()
	logger.Info("Received shutdown signal", "signal", sig.String())

	if err := shutdownFunc(ctx, logger, mgr); err != nil {
		return fmt.Errorf("Shutdown completed with error: %w", err)
	} else {
		logger.Info("Shutdown completed successfully")
	}

	return nil
}
