package mtlsClient

import (
	"net/http"

	baseClient "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr/base"
)

type MtlsKaspiClient struct {
	*baseClient.BaseKaspiClient
}

func New(baseURL string, httpClient *http.Client) *MtlsKaspiClient {
	base := baseClient.NewBaseKaspiClient(baseURL, httpClient)
	return &MtlsKaspiClient{
		BaseKaspiClient: base,
	}
}
