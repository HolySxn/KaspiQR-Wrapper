package ipbasedClient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/HolySxn/KaspiQR-Wrapper/internal/models"
	"github.com/HolySxn/KaspiQR-Wrapper/internal/utils"
)

func (c *IPBasedKaspiClient) GetClientInfo(ctx context.Context, phoneNumber, deviceToken string) (models.ClientInfo, error) {
	url := fmt.Sprintf("%s/remote/client-info?phoneNumber=%s&deviceToken=%s",
		c.BaseURL, phoneNumber, deviceToken)

	data, err := c.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return models.ClientInfo{}, fmt.Errorf("failed to fetch client info: %w", err)
	}

	clientInfo, err := utils.Convert[models.ClientInfo](data)
	if err != nil {
		return models.ClientInfo{}, fmt.Errorf("failed to convert data to ClientInfo: %w", err)
	}

	return clientInfo, nil
}

type RemotePaymentRequest struct {
	OrganizationBin string  `json:"OrganizationBin"`
	Amount          float64 `json:"Amount"`
	PhoneNumber     string  `json:"PhoneNumber"`
	DeviceToken     string  `json:"DeviceToken"`
	Comment         string  `json:"Comment,omitempty"`
}

func (c *IPBasedKaspiClient) CreateRemotePayment(ctx context.Context, amount float64, phoneNumber, deviceToken, comment string) (models.RemotePayment, error) {
	url := c.BaseURL + "/remote/create"

	body := RemotePaymentRequest{
		OrganizationBin: c.BIN,
		Amount:          amount,
		PhoneNumber:     phoneNumber,
		DeviceToken:     deviceToken,
		Comment:         comment,
	}

	data, err := c.DoRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return models.RemotePayment{}, fmt.Errorf("failed to create remote payment: %w", err)
	}

	payment, err := utils.Convert[models.RemotePayment](data)
	if err != nil {
		return models.RemotePayment{}, fmt.Errorf("failed to convert data to RemotePayment: %w", err)
	}

	return payment, nil
}

type RemotePaymentCancelRequest struct {
	OrganizationBin string `json:"OrganizationBin"`
	QrPaymentId     int64  `json:"QrPaymentId"`
	DeviceToken     string `json:"DeviceToken"`
}

func (c *IPBasedKaspiClient) CancelRemotePayment(ctx context.Context, qrPaymentID int64, deviceToken string) (models.CancelRemotePaymentStatus, error) {
	url := fmt.Sprintf("%s/remote/cancel", c.BaseURL)

	body := RemotePaymentCancelRequest{
		OrganizationBin: c.BIN,
		QrPaymentId:     qrPaymentID,
		DeviceToken:     deviceToken,
	}

	data, err := c.DoRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return models.CancelRemotePaymentStatus{}, fmt.Errorf("failed to cancel remote payment: %w", err)
	}

	status, err := utils.Convert[models.CancelRemotePaymentStatus](data)
	if err != nil {
		return models.CancelRemotePaymentStatus{}, fmt.Errorf("failed to convert data to PaymentStatus: %w", err)
	}

	return status, nil
}
