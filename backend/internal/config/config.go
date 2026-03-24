package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"nethub-mdm/pkg/logger"

	"github.com/joho/godotenv"
)

type Config struct {
	Env         string
	HTTPPort    int
	DatabaseURL string
}

func LoadEnvVars(log logger.Logger) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Error finding current file path")
	}

	basePath := filepath.Dir(filename)
	log.Infof("Config package path: %s", basePath)

	envFile := os.Getenv("ENV_PATH")
	if envFile == "" {
		log.Info("ENV_PATH is not set, not loading .env file. Using environment variables from the system.")
		return
	}

	log.Infof("Found ENV_PATH='%s', so loading environment variables from this file...", envFile)

	envFullFilePath := filepath.Join(basePath, envFile)
	err := godotenv.Load(envFullFilePath)
	if err != nil {
		cwd, _ := os.Getwd()
		log.Fatalf("Error loading .env file: %v. Current directory '%s', tried path '%s'", err.Error(), cwd, envFullFilePath)
	} else {
		log.Infof("Loaded .env file '%s'", envFullFilePath)
	}
}

func GetEnvAsInt(log logger.Logger, key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Infof("Invalid value for %s: %s. Using default: %d", key, valueStr, defaultValue)
		return defaultValue
	}
	return value
}

func GetEnvAsString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func LoadConfig(log logger.Logger) *Config {
	LoadEnvVars(log)

	return &Config{
		Env:         GetEnvAsString("ENV", "development"),
		HTTPPort:    GetEnvAsInt(log, "PORT", 9000),
		DatabaseURL: GetEnvAsString("DATABASE_URL", "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
	}
}
