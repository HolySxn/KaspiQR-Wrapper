package kaspihandler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/HolySxn/KaspiQR-Wrapper/internal/models"
	"github.com/HolySxn/KaspiQR-Wrapper/internal/utils"
)

func (h *KaspiHandler) HandleGetTradePoints(ctx context.Context) ([]models.TradePoint, error) {
	url := h.baseURL + "/partner/tradepoints"

	data, err := h.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	tradePoints, err := utils.Convert[[]models.TradePoint](data)
	if err != nil {
		return nil, fmt.Errorf("failed to convert data to trade points: %w", err)
	}

	return tradePoints, nil
}

type deviceRegisterRequest struct {
	DeviceID     string `json:"deviceId"`
	TradePointID int64  `json:"tradePointId"`
}

func (h *KaspiHandler) HandleDeviceRegister(ctx context.Context, deviceID string, tradePointID int64) (models.DeviceToken, error) {
	url := h.baseURL + "/device/register"
	body := deviceRegisterRequest{
		DeviceID:     deviceID,
		TradePointID: tradePointID,
	}

	data, err := h.doRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return models.DeviceToken{}, err
	}

	deviceToken, err := utils.Convert[models.DeviceToken](data)
	if err != nil {
		return models.DeviceToken{}, fmt.Errorf("failed to convert data to deviceToken points: %w", err)
	}

	return deviceToken, nil
}
