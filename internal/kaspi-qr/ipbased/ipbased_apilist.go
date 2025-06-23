package ipbasedClient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/HolySxn/KaspiQR-Wrapper/internal/models"
	"github.com/HolySxn/KaspiQR-Wrapper/internal/utils"
)

func (c *IPBasedKaspiClient) GetTradePointsByBin(ctx context.Context) ([]models.TradePoint, error) {
	url := fmt.Sprintf("%s/partner/tradepoints/%s", c.BaseURL, c.BIN)

	data, err := c.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch trade points for BIN %s: %w", c.BIN, err)
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
	BIN          string `json:"OrganizationBin"`
}

func (c *IPBasedKaspiClient) DeviceRegister(ctx context.Context, deviceID string, tradePointID int64) (models.DeviceToken, error) {
	url := c.BaseURL + "/device/register"
	body := deviceRegisterRequest{
		DeviceID:     deviceID,
		TradePointID: tradePointID,
		BIN:          c.BIN,
	}

	data, err := c.DoRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return models.DeviceToken{}, err
	}

	deviceToken, err := utils.Convert[models.DeviceToken](data)
	if err != nil {
		return models.DeviceToken{}, fmt.Errorf("failed to convert data to deviceToken: %w", err)
	}

	return deviceToken, nil
}

type deviceDeleteRequest struct {
	DeviceToken string `json:"DeviceToken"`
	BIN         string `json:"OrganizationBin"`
}

func (c *IPBasedKaspiClient) DeviceDelete(ctx context.Context, deviceToken string) error {
	url := c.BaseURL + "/device/delete"
	body := deviceDeleteRequest{
		DeviceToken: deviceToken,
		BIN:         c.BIN,
	}

	_, err := c.DoRequest(ctx, http.MethodPost, url, body)
	return err
}
