package ipbasedClient

import (
	"net/http"

	baseClient "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr/base"
)

type APIKeyKaspiClient struct {
	*baseClient.BaseKaspiClient
	BIN string
}

func New(baseURL, bin string, httpClient *http.Client) *APIKeyKaspiClient {
	return &APIKeyKaspiClient{
		BaseKaspiClient: &baseClient.BaseKaspiClient{
			BaseURL: baseURL,
			Client:  httpClient,
		},
		BIN: bin,
	}
}
