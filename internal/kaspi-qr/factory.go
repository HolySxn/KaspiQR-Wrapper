package kaspiqr

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/HolySxn/KaspiQR-Wrapper/config"
	apikeyclient "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr/apikey"
	ipbasedClient "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr/ipbased"
	mtlsClient "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr/mtls"
)

func NewKaspiClient(cfg *config.Config) (KaspiQRBase, error) {
	var httpClient *http.Client
	var tlsConfig *tls.Config
	var err error

	switch cfg.Kaspi.AuthMode {
	case config.AuthModeAPIKey:
		tlsConfig, err = tlsConfigWithRootCA(cfg.Kaspi.CACert)
	case config.AuthModeMTLS, config.AuthModeIPBased:
		tlsConfig, err = tlsConfigWithClientCert(cfg.Kaspi.ClientCert, cfg.Kaspi.ClientKey, cfg.Kaspi.CACert)
	default:
		return nil, fmt.Errorf("unsupported authentication mode: %s", cfg.Kaspi.AuthMode)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create TLS config: %w", err)
	}

	httpClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	switch cfg.Kaspi.AuthMode {
	case config.AuthModeAPIKey:
		return apikeyclient.New(cfg.Kaspi.BaseURL, cfg.Kaspi.APIKey, httpClient), nil
	case config.AuthModeMTLS:
		return mtlsClient.New(cfg.Kaspi.BaseURL, httpClient), nil
	case config.AuthModeIPBased:
		return ipbasedClient.New(cfg.Kaspi.BaseURL, cfg.Kaspi.CompanyBIN, httpClient), nil
	}

	return nil, fmt.Errorf("unexpected error initializing Kaspi client")
}

func tlsConfigWithClientCert(certFile, keyFile, caFile string) (*tls.Config, error) {
	config, err := tlsConfigWithRootCA(caFile)
	if err != nil {
		return nil, err
	}

	clientCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	config.Certificates = []tls.Certificate{clientCert}

	return config, nil
}

func tlsConfigWithRootCA(caFile string) (*tls.Config, error) {
	caCertPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	caCert, err := os.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load CA cert: %w", err)
	}

	if !caCertPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to append CA cert to pool")
	}

	config := &tls.Config{
		RootCAs:    caCertPool,
		MinVersion: tls.VersionTLS12,
	}

	return config, nil
}
