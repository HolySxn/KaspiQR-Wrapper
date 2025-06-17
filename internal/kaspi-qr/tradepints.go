package kaspihandler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/HolySxn/KaspiQR-Wrapper/internal/models"
	"github.com/HolySxn/KaspiQR-Wrapper/internal/utils"
	"github.com/google/uuid"
)

type KaspiHandler struct {
	Client  *http.Client
	baseURL string
	apiKey  string
}

func NewKaspiHandler(baseURL string, apiKey string) *KaspiHandler {
	client := &http.Client{
		Timeout: time.Second,
	}

	return &KaspiHandler{
		Client:  client,
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

func (h *KaspiHandler) HandleGetTradePoints(ctx context.Context) ([]models.TradePoint, error) {
	url := h.baseURL + "/partner/tradepoints"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Request-ID", uuid.New().String())
	if h.apiKey != "" {
		req.Header.Set("Api-Key", h.apiKey)
	}

	res, err := h.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	var rawResp models.RawResponse
	if err := json.NewDecoder(res.Body).Decode(&rawResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	tradePoints, err := utils.Convert[[]models.TradePoint](rawResp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to convert data to trade points: %w", err)
	}

	return tradePoints, nil
}
