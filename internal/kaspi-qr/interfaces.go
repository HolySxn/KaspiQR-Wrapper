package kaspiqr

import (
	"context"

	"github.com/HolySxn/KaspiQR-Wrapper/internal/models"
)

type KaspiQR interface {
	GetTradePoints(ctx context.Context) ([]models.TradePoint, error)
	DeviceRegister(ctx context.Context, deviceID string, tradePointID int64) (models.DeviceToken, error)
	DeviceDelete(ctx context.Context, deviceToken string) error
	CreateQR(ctx context.Context, deviceToken string, amount float64, externalID string) (models.QrToken, error)
	CreateLink(ctx context.Context, deviceToken string, amount float64, externalID string) (models.PaymentData, error)
	GetPaymentStatus(ctx context.Context, qrPaymentToken string) (models.PaymentStatus, error)
}
