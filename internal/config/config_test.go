package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	// Save original environment variables
	originalEnv := make(map[string]string)
	envVars := []string{"PORT", "LOG_LEVEL", "READ_TIMEOUT", "WRITE_TIMEOUT", "ENVIRONMENT"}

	for _, env := range envVars {
		originalEnv[env] = os.Getenv(env)
	}

	// Clean up after test
	defer func() {
		for _, env := range envVars {
			if val, exists := originalEnv[env]; exists {
				os.Setenv(env, val)
			} else {
				os.Unsetenv(env)
			}
		}
	}()

	t.Run("default values", func(t *testing.T) {
		// Unset all environment variables
		for _, env := range envVars {
			os.Unsetenv(env)
		}

		cfg := Load()

		assert.Equal(t, 8080, cfg.Port)
		assert.Equal(t, "info", cfg.LogLevel)
		assert.Equal(t, 10, cfg.ReadTimeout)
		assert.Equal(t, 10, cfg.WriteTimeout)
		assert.Equal(t, "development", cfg.Environment)
	})

	t.Run("custom values", func(t *testing.T) {
		os.Setenv("PORT", "3000")
		os.Setenv("LOG_LEVEL", "debug")
		os.Setenv("READ_TIMEOUT", "30")
		os.Setenv("WRITE_TIMEOUT", "30")
		os.Setenv("ENVIRONMENT", "production")

		cfg := Load()

		assert.Equal(t, 3000, cfg.Port)
		assert.Equal(t, "debug", cfg.LogLevel)
		assert.Equal(t, 30, cfg.ReadTimeout)
		assert.Equal(t, 30, cfg.WriteTimeout)
		assert.Equal(t, "production", cfg.Environment)
	})

	t.Run("invalid port", func(t *testing.T) {
		os.Setenv("PORT", "invalid")

		cfg := Load()

		// Should fall back to default
		assert.Equal(t, 8080, cfg.Port)
	})
}

func TestGetEnv(t *testing.T) {
	t.Run("existing environment variable", func(t *testing.T) {
		os.Setenv("TEST_VAR", "test_value")
		defer os.Unsetenv("TEST_VAR")

		result := getEnv("TEST_VAR", "default")
		assert.Equal(t, "test_value", result)
	})

	t.Run("non-existing environment variable", func(t *testing.T) {
		os.Unsetenv("NON_EXISTING_VAR")

		result := getEnv("NON_EXISTING_VAR", "default")
		assert.Equal(t, "default", result)
	})
}

func TestGetEnvAsInt(t *testing.T) {
	tests := []struct {
		name         string
		envValue     string
		defaultValue int
		expected     int
	}{
		{
			name:         "valid integer",
			envValue:     "3000",
			defaultValue: 8080,
			expected:     3000,
		},
		{
			name:         "invalid integer",
			envValue:     "invalid",
			defaultValue: 8080,
			expected:     8080,
		},
		{
			name:         "empty value",
			envValue:     "",
			defaultValue: 8080,
			expected:     8080,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv("TEST_INT_VAR", tt.envValue)
				defer os.Unsetenv("TEST_INT_VAR")
			} else {
				os.Unsetenv("TEST_INT_VAR")
			}

			result := getEnvAsInt("TEST_INT_VAR", tt.defaultValue)
			assert.Equal(t, tt.expected, result)
		})
	}
}
