package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("KASPI_API_KEY", "test_api_key")
	os.Setenv("KASPI_CLIENT_CERT", "test_client_cert")
	os.Setenv("KASPI_CLIENT_KEY", "test_client_key")
	os.Setenv("KASPI_CA_CERT", "test_ca_cert")
	defer func() {
		os.Unsetenv("KASPI_API_KEY")
		os.Unsetenv("KASPI_CLIENT_CERT")
		os.Unsetenv("KASPI_CLIENT_KEY")
		os.Unsetenv("KASPI_CA_CERT")
	}()

	configContent := `
server:
  port: "8080"
  host: "localhost"
kaspi:
  base_url: "https://kaspi.kz"
  auth_mode: "api_key"
`
	configPath := "test_config.yaml"
	err := os.WriteFile(configPath, []byte(configContent), 0644)
	assert.NoError(t, err)
	defer os.Remove(configPath)

	cfg, err := LoadConfig(configPath)
	assert.NoError(t, err)

	assert.Equal(t, "8080", cfg.Server.Port)
	assert.Equal(t, "localhost", cfg.Server.Host)
	assert.Equal(t, "https://kaspi.kz", cfg.Kaspi.BaseURL)
	assert.Equal(t, AuthModeAPIKey, cfg.Kaspi.AuthMode)
	assert.Equal(t, "test_api_key", cfg.Kaspi.APIKey)
	assert.Equal(t, "test_client_cert", cfg.Kaspi.ClientCert)
	assert.Equal(t, "test_client_key", cfg.Kaspi.ClientKey)
	assert.Equal(t, "test_ca_cert", cfg.Kaspi.CACert)
}
