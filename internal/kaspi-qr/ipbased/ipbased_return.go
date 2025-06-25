package ipbasedClient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/HolySxn/KaspiQR-Wrapper/internal/models"
	"github.com/HolySxn/KaspiQR-Wrapper/internal/utils"
)

type PaymentReturnRequest struct {
	DeviceToken string  `json:"DeviceToken"`
	QrPaymentId int64   `json:"QrPaymentId"`
	Amount      float64 `json:"Amount"`
	BIN         string  `json:"OrganizationBin"`
}

func (c *IPBasedKaspiClient) PaymentReturn(ctx context.Context, deviceToken string, qrPaymentID int64, amount float64) (models.ReturnOperationId, error) {
	url := c.BaseURL + "/payment/return"

	body := PaymentReturnRequest{
		DeviceToken: deviceToken,
		QrPaymentId: qrPaymentID,
		Amount:      amount,
		BIN:         c.BIN,
	}

	data, err := c.DoRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return models.ReturnOperationId{}, err
	}

	returrnId, err := utils.Convert[models.ReturnOperationId](data)
	if err != nil {
		return models.ReturnOperationId{}, fmt.Errorf("failed to convert data to ReturnOperationId: %w", err)
	}

	return returrnId, nil
}
