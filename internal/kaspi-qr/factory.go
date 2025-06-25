package kaspiqr

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"

	"github.com/HolySxn/KaspiQR-Wrapper/config"
	apikeyclient "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr/apikey"
	ipbasedClient "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr/ipbased"
	mtlsClient "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr/mtls"
)

func NewKaspiClient(cfg *config.Config) (KaspiQRBase, error) {
	var httpClient *http.Client

	switch cfg.Kaspi.AuthMode {
	case config.AuthModeAPIKey:
		httpClient = &http.Client{}
	case config.AuthModeMTLS, config.AuthModeIPBased:
		tlsConfig, err := tlsConfig(cfg.Kaspi.ClientCert, cfg.Kaspi.ClientKey, cfg.Kaspi.CACert)
		if err != nil {
			return nil, err
		}

		httpClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		}
	default:
		return nil, fmt.Errorf("unsupported authentication mode: %s", cfg.Kaspi.AuthMode)
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

func tlsConfig(certFile, keyFile, caFile string) (*tls.Config, error) {
	clientCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	// caCert, err := os.ReadFile(caFile)
	// if err != nil {
	// 	return nil, err
	// }

	// caCertPool := x509.NewCertPool()
	// if !caCertPool.AppendCertsFromPEM(caCert) {
	// 	return nil, fmt.Errorf("failed to append CA cert")
	// }

	caCertPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      caCertPool,
		MinVersion:   tls.VersionTLS12,
	}

	return config, nil
}
