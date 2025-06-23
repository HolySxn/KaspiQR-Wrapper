package mtlsClient

import (
	"context"
	"fmt"
	"net/http"

	baseClient "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr/base"
	"github.com/HolySxn/KaspiQR-Wrapper/internal/models"
	"github.com/HolySxn/KaspiQR-Wrapper/internal/utils"
)

type MtlsKaspiClient struct {
	*baseClient.BaseKaspiClient
}

func New(baseURL string, httpClient *http.Client) *MtlsKaspiClient {
	return &MtlsKaspiClient{
		BaseKaspiClient: &baseClient.BaseKaspiClient{
			BaseURL: baseURL,
			Client:  httpClient,
		},
	}
}

type CreateReturnRequest struct {
	DeviceToken string `json:"DeviceToken"`
	ExternalID  string `json:"ExternalId"`
}

func (c *MtlsKaspiClient) CreateReturn(ctx context.Context, deviceToken string, externalID string) (models.Return, error) {
	url := c.BaseURL + "/return/create"
	body := CreateReturnRequest{
		DeviceToken: deviceToken,
		ExternalID:  externalID,
	}

	data, err := c.DoRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return models.Return{}, err
	}

	status, err := utils.Convert[models.Return](data)
	if err != nil {
		return models.Return{}, fmt.Errorf("failed to convert data to Return response: %w", err)
	}

	return status, nil
}

func (c *MtlsKaspiClient) GetReturnStatus(ctx context.Context, qrReturnID int64) (models.ReturnStatus, error) {
	url := c.BaseURL + "/return/status/" + fmt.Sprintf("%d", qrReturnID)

	data, err := c.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return models.ReturnStatus{}, err
	}

	status, err := utils.Convert[models.ReturnStatus](data)
	if err != nil {
		return models.ReturnStatus{}, fmt.Errorf("failed to convert data to Return status: %w", err)
	}

	return status, nil
}

type RecentOperationsRequest struct {
	DeviceToken string `json:"DeviceToken"`
	QrReturnId  int64  `json:"QrReturnId"`
	MaxResult   int64  `json:"MaxResult"`
}

func (c *MtlsKaspiClient) ReturnOperations(ctx context.Context, deviceToken string, qrReturnID int64, maxResult int64) ([]models.RecentOperation, error) {
	url := c.BaseURL + "/return/operations"
	body := RecentOperationsRequest{
		DeviceToken: deviceToken,
		QrReturnId:  qrReturnID,
		MaxResult:   maxResult,
	}

	data, err := c.DoRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	operations, err := utils.Convert[[]models.RecentOperation](data)
	if err != nil {
		return nil, fmt.Errorf("failed to convert data to RecentOperations: %w", err)
	}

	return operations, nil
}

type PaymentDetailsRequest struct {
	QrPaymentId int64  `json:"QrPaymentId"`
	DeviceToken string `json:"DeviceToken"`
}

func (c *MtlsKaspiClient) PaymentDetails(ctx context.Context, qrPaymentID int64, deviceToken string) (models.PaymentDetails, error) {
	url := c.BaseURL + "/payment/details"

	body := PaymentDetailsRequest{
		QrPaymentId: qrPaymentID,
		DeviceToken: deviceToken,
	}

	data, err := c.DoRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return models.PaymentDetails{}, err
	}

	details, err := utils.Convert[models.PaymentDetails](data)
	if err != nil {
		return models.PaymentDetails{}, fmt.Errorf("failed to convert data to PaymentDetails: %w", err)
	}

	return details, nil
}

type PaymentReturnRequest struct {
	DeviceToken string  `json:"DeviceToken"`
	QrPaymentId int64   `json:"QrPaymentId"`
	QrReturnId  int64   `json:"QrReturnId"`
	Amount      float64 `json:"Amount"`
}

func (c *MtlsKaspiClient) PaymentReturn(ctx context.Context, deviceToken string, qrPaymentID int64, qrReturnID int64, amount float64) (models.ReturnOperationId, error) {
	url := c.BaseURL + "/payment/return"

	body := PaymentReturnRequest{
		DeviceToken: deviceToken,
		QrPaymentId: qrPaymentID,
		QrReturnId:  qrReturnID,
		Amount:      amount,
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
