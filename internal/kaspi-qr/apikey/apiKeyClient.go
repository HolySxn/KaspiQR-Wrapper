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
	base := baseClient.NewBaseKaspiClient(baseURL, httpClient)
	apiClient := &APIKeyKaspiClient{
		BaseKaspiClient: base,
		APIKey:          apiKey,
	}
	apiClient.HeadSetter = apiClient.setHeader
	return apiClient
}

func (c *APIKeyKaspiClient) setHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Request-ID", uuid.New().String())
	req.Header.Set("Api-Key", c.APIKey)
}
