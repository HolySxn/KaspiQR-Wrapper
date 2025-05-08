package apikeyclient

import (
	"net/http"

	baseClient "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr/base"
	"github.com/google/uuid"
)

type APIKeyKaspiClient struct {
	*baseClient.BaseKaspiClient
	APIKey string
}

func New(baseURL, apiKey string, httpClient *http.Client) *APIKeyKaspiClient {
	return &APIKeyKaspiClient{
		BaseKaspiClient: &baseClient.BaseKaspiClient{
			BaseURL: baseURL,
			Client:  httpClient,
		},
		APIKey: apiKey,
	}
}

func (c *APIKeyKaspiClient) SeteHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Request-ID", uuid.New().String())
	req.Header.Set("Api-Key", c.APIKey)
}
