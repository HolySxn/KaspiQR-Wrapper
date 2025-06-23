package base

import (
	"context"
	"fmt"
	"net/http"

	"github.com/HolySxn/KaspiQR-Wrapper/internal/models"
	"github.com/HolySxn/KaspiQR-Wrapper/internal/utils"
)

func (c *BaseKaspiClient) GetTradePoints(ctx context.Context) ([]models.TradePoint, error) {
	url := c.BaseURL + "/partner/tradepoints"

	data, err := c.DoRequest(ctx, http.MethodGet, url, nil)
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

func (c *BaseKaspiClient) DeviceRegister(ctx context.Context, deviceID string, tradePointID int64) (models.DeviceToken, error) {
	url := c.BaseURL + "/device/register"
	body := deviceRegisterRequest{
		DeviceID:     deviceID,
		TradePointID: tradePointID,
	}

	data, err := c.DoRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return models.DeviceToken{}, err
	}

	deviceToken, err := utils.Convert[models.DeviceToken](data)
	if err != nil {
		return models.DeviceToken{}, fmt.Errorf("failed to convert data to deviceToken points: %w", err)
	}

	return deviceToken, nil
}

type deviceDeleteRequest struct {
	DeviceToken string `json:"DeviceToken"`
}

func (c *BaseKaspiClient) DeviceDelete(ctx context.Context, deviceToken string) error {
	url := c.BaseURL + "/device/delete"
	body := deviceDeleteRequest{
		DeviceToken: deviceToken,
	}

	_, err := c.DoRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return err
	}

	return nil
}
