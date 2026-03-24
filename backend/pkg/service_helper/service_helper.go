package service_helper

import (
	"context"
	"fmt"
	"nethub-mdm/internal/storage"
	"nethub-mdm/pkg/db_manager"
	"nethub-mdm/pkg/logger"
	"nethub-mdm/pkg/shutdown"
	"os"
)

type LifecycleFunc func(ctx context.Context, log logger.Logger, mgr db_manager.Manager) error

func StartService(
	serviceName string,
	logger logger.Logger,
	dsn string,
	initFunc LifecycleFunc,
	startFunc LifecycleFunc,
	shutdownFunc LifecycleFunc,
) {
	prefix := fmt.Sprintf("Service '%s'", serviceName)
	logger.Infof("%s: Initializing", prefix)
	defer logger.Infof("%s: Exited", prefix)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mgr, err := db_manager.NewDbManager(dsn)
	if err != nil {
		logger.Error("Error connecting to database: %v", err)
		os.Exit(1)
	}

	err = mgr.UsePlugins(
		storage.NewAuditPlugin(logger),
	)
	if err != nil {
		logger.Error("initialization plugins failed", "error", err)
	}

	logger.Info("Starting initialization...")

	if err := initFunc(ctx, logger, mgr); err != nil {
		logger.Error("Initialization failed", "error", err)
		os.Exit(1)
	}

	if err := startFunc(ctx, logger, mgr); err != nil {
		logger.Error("Start failed", "error", err)
		os.Exit(1)
	}

	logger.Info("Service is running. Waiting for shutdown signal...")
	sig := shutdown.WaitForSignalToShutdown()
	logger.Info("Received shutdown signal", "signal", sig.String())

	if err := shutdownFunc(ctx, logger, mgr); err != nil {
		logger.Error("Shutdown completed with error", "error", err)
		os.Exit(1)
	} else {
		logger.Info("Shutdown completed successfully")
	}
}
