package kaspihandler

import (
	"log/slog"
	"net/http"
	"time"
)

type KaspiHandler struct {
	Client  *http.Client
	logger  *slog.Logger
	baseURL string
}

func NewKaspiHandler(baseURL string, logger *slog.Logger) *KaspiHandler {
	client := &http.Client{
		Timeout: time.Duration(1) * time.Second,
	}

	return &KaspiHandler{
		Client:  client,
		logger:  logger,
		baseURL: baseURL,
	}
}

// func (h *KaspiHandler) HandleGetTradePoints(ctx context.Context) ([]models.TradePoint, error) {
// 	url := h.baseURL + "/partner/tradepoints"
// 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create request: %w", err)
// 	}

// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Accept", "application/json")
// 	req.Header.Set("X-Request-ID", uuid.New().String())

// 	res, err := h.Client.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("request failed: %w", err)
// 	}
// }
