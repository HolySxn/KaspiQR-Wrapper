package ipbasedClient

import (
	"net/http"

	baseClient "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr/base"
)

type IPBasedKaspiClient struct {
	*baseClient.BaseKaspiClient
	BIN string
}

func New(baseURL, bin string, httpClient *http.Client) *IPBasedKaspiClient {
	base := baseClient.NewBaseKaspiClient(baseURL, httpClient)
	return &IPBasedKaspiClient{
		BaseKaspiClient: base,
		BIN:             bin,
	}
}
