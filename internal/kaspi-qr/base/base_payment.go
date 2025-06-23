package base

import (
	"context"
	"fmt"
	"net/http"

	"github.com/HolySxn/KaspiQR-Wrapper/internal/models"
	"github.com/HolySxn/KaspiQR-Wrapper/internal/utils"
)

type CreateRequest struct {
	DeviceToken string  `json:"DeviceToken"`
	Amount      float64 `json:"Amount"`
	ExternalID  string  `json:"ExternalId"`
}

func (c *BaseKaspiClient) CreateQR(ctx context.Context, deviceToken string, amount float64, externalID string) (models.QrToken, error) {
	url := c.BaseURL + "/qr/create"
	body := CreateRequest{
		DeviceToken: deviceToken,
		Amount:      amount,
		ExternalID:  externalID,
	}

	data, err := c.DoRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return models.QrToken{}, err
	}

	qr, err := utils.Convert[models.QrToken](data)
	if err != nil {
		return models.QrToken{}, fmt.Errorf("failed to convert data to qrToken points: %w", err)
	}

	return qr, nil
}

func (c *BaseKaspiClient) CreateLink(ctx context.Context, deviceToken string, amount float64, externalID string) (models.PaymentData, error) {
	url := c.BaseURL + "/qr/create-link"
	body := CreateRequest{
		DeviceToken: deviceToken,
		Amount:      amount,
		ExternalID:  externalID,
	}

	data, err := c.DoRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return models.PaymentData{}, err
	}

	link, err := utils.Convert[models.PaymentData](data)
	if err != nil {
		return models.PaymentData{}, fmt.Errorf("failed to convert data to PaymentData: %w", err)
	}

	return link, nil
}

func (c *BaseKaspiClient) GetPaymentStatus(ctx context.Context, qrPaymentId int64) (models.PaymentStatus, error) {
	url := fmt.Sprintf("%s/payment/status/%d", c.BaseURL, qrPaymentId)

	data, err := c.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return models.PaymentStatus{}, err
	}

	status, err := utils.Convert[models.PaymentStatus](data)
	if err != nil {
		return models.PaymentStatus{}, fmt.Errorf("failed to convert data to PaymentStatus: %w", err)
	}

	return status, nil
}
