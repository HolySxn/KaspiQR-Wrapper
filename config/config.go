package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

const (
	AuthModeAPIKey  string = "api_key"
	AuthModeMTLS    string = "mtls"
	AuthModeIPBased string = "ip_based"
)

type ServerConfig struct {
	Port   string `env:"PORT,default=8000"`
	Host   string `env:"HOST,default=0.0.0.0"`
	LogLvl string `env:"LOG_LEVEL,default=debug"`
}

type KaspiConfig struct {
	BaseURL    string `env:"BASE_URL"`
	AuthMode   string `env:"AUTH_MODE"`
	APIKey     string `env:"API_KEY"`
	CompanyBIN string `env:"COMPANY_BIN"`
	ClientCert string `env:"CLIENT_CERT"`
	ClientKey  string `env:"CLIENT_KEY"`
	CACert     string `env:"CA_CERT"`
}

type PostgresConfig struct {
	Host     string `env:"HOST,default=localhost"`
	Port     string `env:"PORT,default=5432"`
	User     string `env:"USER,default=postgres"`
	Password string `env:"PASSWORD"`
	DBName   string `env:"DB,default=postgres"`
	SSLMode  string `env:"SSLMODE,default=disable"`
}

type Config struct {
	Server   ServerConfig   `env:",prefix=SERVER_"`
	Kaspi    KaspiConfig    `env:",prefix=KASPI_"`
	Postgres PostgresConfig `env:",prefix=POSTGRES_"`
}

func LoadConfig(ctx context.Context) (*Config, error) {
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
