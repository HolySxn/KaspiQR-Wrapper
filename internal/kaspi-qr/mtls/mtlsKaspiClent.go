package mtlsClient

import (
	"net/http"

	baseClient "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr/base"
)

type APIKeyKaspiClient struct {
	*baseClient.BaseKaspiClient
}

func New(baseURL string, httpClient *http.Client) *APIKeyKaspiClient {
	return &APIKeyKaspiClient{
		BaseKaspiClient: &baseClient.BaseKaspiClient{
			BaseURL: baseURL,
			Client:  httpClient,
		},
	}
}
