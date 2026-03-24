package config

import (
	"nethub-mdm/pkg/logger"
	"os"
	"testing"
)

func TestGetEnvAsInt(t *testing.T) {
	log, _ := logger.NewZapLogger("test")

	t.Run("existing env", func(t *testing.T) {
		os.Setenv("TEST_PORT", "8080")
		defer os.Unsetenv("TEST_PORT")

		val := GetEnvAsInt(log, "TEST_PORT", 9000)
		if val != 8080 {
			t.Errorf("expected 8080, got %d", val)
		}
	})

	t.Run("default value", func(t *testing.T) {
		val := GetEnvAsInt(log, "NON_EXISTENT", 9000)
		if val != 9000 {
			t.Errorf("expected 9000, got %d", val)
		}
	})

	t.Run("invalid value", func(t *testing.T) {
		os.Setenv("TEST_BAD", "abc")
		defer os.Unsetenv("TEST_BAD")

		val := GetEnvAsInt(log, "TEST_BAD", 9000)
		if val != 9000 {
			t.Errorf("expected default 9000 for bad input, got %d", val)
		}
	})
}
