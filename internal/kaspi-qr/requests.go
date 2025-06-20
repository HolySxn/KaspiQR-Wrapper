package kaspihandler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/HolySxn/KaspiQR-Wrapper/internal/models"
	"github.com/HolySxn/KaspiQR-Wrapper/internal/utils"
)

func (h *KaspiHandler) GetTradePoints(ctx context.Context) ([]models.TradePoint, error) {
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

func (h *KaspiHandler) DeviceRegister(ctx context.Context, deviceID string, tradePointID int64) (models.DeviceToken, error) {
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

type deviceDeleteRequest struct {
	DeviceToken string `json:"DeviceToken"`
}

func (h *KaspiHandler) DeviceDelete(ctx context.Context, deviceToken string) error {
	url := h.baseURL + "/device/delete"
	body := deviceDeleteRequest{
		DeviceToken: deviceToken,
	}

	_, err := h.doRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return err
	}

	return nil
}

type CreateRequest struct {
	DeviceToken string  `json:"DeviceToken"`
	Amount      float64 `json:"Amount"`
	ExternalID  string  `json:"ExternalId"`
}

func (h *KaspiHandler) CreateQR(ctx context.Context, deviceToken string, amount float64, externalID string) (models.QrToken, error) {
	url := h.baseURL + "/qr/create"
	body := CreateRequest{
		DeviceToken: deviceToken,
		Amount:      amount,
		ExternalID:  externalID,
	}

	data, err := h.doRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return models.QrToken{}, err
	}

	qr, err := utils.Convert[models.QrToken](data)
	if err != nil {
		return models.QrToken{}, fmt.Errorf("failed to convert data to qrToken points: %w", err)
	}

	return qr, nil
}

func (h *KaspiHandler) CreateLink(ctx context.Context, deviceToken string, amount float64, externalID string) (models.PaymentData, error) {
	url := h.baseURL + "/qr/create-link"
	body := CreateRequest{
		DeviceToken: deviceToken,
		Amount:      amount,
		ExternalID:  externalID,
	}

	data, err := h.doRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return models.PaymentData{}, err
	}

	link, err := utils.Convert[models.PaymentData](data)
	if err != nil {
		return models.PaymentData{}, fmt.Errorf("failed to convert data to PaymentData: %w", err)
	}

	return link, nil
}

func (h *KaspiHandler) GetPaymentStatus(ctx context.Context, deviceToken string, qrPaymentToken string) (models.PaymentStatus, error) {
	url := h.baseURL + "/payment/status/" + qrPaymentToken

	data, err := h.doRequest(ctx, http.MethodPost, url, nil)
	if err != nil {
		return models.PaymentStatus{}, err
	}

	status, err := utils.Convert[models.PaymentStatus](data)
	if err != nil {
		return models.PaymentStatus{}, fmt.Errorf("failed to convert data to PaymentStatus: %w", err)
	}

	return status, nil
}