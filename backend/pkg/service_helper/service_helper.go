package service_helper

import (
	"context"
	"nethub-mdm/pkg/logger"
	"nethub-mdm/pkg/shutdown"
	"os"
)

type LifecycleFunc func(ctx context.Context, log logger.Logger) error

func StartService(
	serviceName string,
	initFunc LifecycleFunc,
	startFunc LifecycleFunc,
	shutdownFunc LifecycleFunc,
) {
	log, err := logger.NewZapLogger(serviceName)
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer log.Sync()

	ctx := context.Background()

	log.Info("Starting initialization...")

	if err := initFunc(ctx, log); err != nil {
		log.Error("Initialization failed", "error", err)
		os.Exit(1)
	}

	if err := startFunc(ctx, log); err != nil {
		log.Error("Start failed", "error", err)
		os.Exit(1)
	}

	log.Info("Service is running. Waiting for shutdown signal...")
	sig := shutdown.WaitForSignalToShutdown()
	log.Info("Received shutdown signal", "signal", sig.String())

	if err := shutdownFunc(ctx, log); err != nil {
		log.Error("Shutdown completed with error", "error", err)
		os.Exit(1)
	} else {
		log.Info("Shutdown completed successfully")
	}
}
