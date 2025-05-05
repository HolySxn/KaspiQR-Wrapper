package config

import (
	"log/slog"
	"os"

	"github.com/spf13/viper"
)

type AuthMode string

const (
	AuthModeAPIKey  AuthMode = "api_key"
	AuthModeMTLS    AuthMode = "mtls"
	AuthModeIPBased AuthMode = "ip_based"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"server"`
	Kaspi struct {
		BaseURL  string   `mapstructure:"base_url"`
		AuthMode AuthMode `mapstructure:"auth_mode"`

		APIKey     string `mapstructure:"-"`
		ClientCert string `mapstructure:"-"`
		ClientKey  string `mapstructure:"-"`
		CACert     string `mapstructure:"-"`
	} `mapstructure:"kaspi"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	cfg.Kaspi.APIKey = os.Getenv("KASPI_API_KEY")
	cfg.Kaspi.ClientCert = os.Getenv("KASPI_CLIENT_CERT")
	cfg.Kaspi.ClientKey = os.Getenv("KASPI_CLIENT_KEY")
	cfg.Kaspi.CACert = os.Getenv("KASPI_CA_CERT")

	return &cfg, nil
}

func NewLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	return logger
}
