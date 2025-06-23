package mtlsClient

import (
	"net/http"

	baseClient "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr/base"
)

type MtlsKaspiClient struct {
	*baseClient.BaseKaspiClient
}

func New(baseURL string, httpClient *http.Client) *MtlsKaspiClient {
	return &MtlsKaspiClient{
		BaseKaspiClient: &baseClient.BaseKaspiClient{
			BaseURL: baseURL,
			Client:  httpClient,
		},
	}
}
