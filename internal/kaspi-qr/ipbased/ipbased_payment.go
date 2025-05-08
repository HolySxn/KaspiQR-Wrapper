package ipbasedClient

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
	BIN         string  `json:"OrganizationBin"`
}

func (c *IPBasedKaspiClient) CreateQR(ctx context.Context, deviceToken string, amount float64, externalID string) (models.QrToken, error) {
	url := c.BaseURL + "/qr/create"
	body := CreateRequest{
		DeviceToken: deviceToken,
		Amount:      amount,
		ExternalID:  externalID,
		BIN:         c.BIN,
	}

	data, err := c.DoRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return models.QrToken{}, err
	}

	qr, err := utils.Convert[models.QrToken](data)
	if err != nil {
		return models.QrToken{}, fmt.Errorf("failed to convert data to qrToken: %w", err)
	}

	return qr, nil
}

func (c *IPBasedKaspiClient) CreateLink(ctx context.Context, deviceToken string, amount float64, externalID string) (models.PaymentData, error) {
	url := c.BaseURL + "/qr/create-link"
	body := CreateRequest{
		DeviceToken: deviceToken,
		Amount:      amount,
		ExternalID:  externalID,
		BIN:         c.BIN,
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

func (c *IPBasedKaspiClient) GetPaymentDetails(ctx context.Context, qrPaymentID string, deviceToken string) (models.PaymentDetails, error) {
	url := fmt.Sprintf("%s/payment/details?QrPaymentId=%s&DeviceToken=%s", c.BaseURL, qrPaymentID, deviceToken)

	data, err := c.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return models.PaymentDetails{}, fmt.Errorf("failed to fetch payment details: %w", err)
	}

	paymentDetails, err := utils.Convert[models.PaymentDetails](data)
	if err != nil {
		return models.PaymentDetails{}, fmt.Errorf("failed to convert data to PaymentDetails: %w", err)
	}

	return paymentDetails, nil
}