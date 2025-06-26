package kaspiqr

import (
	"context"

	mtlsClient "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr/mtls"
	"github.com/HolySxn/KaspiQR-Wrapper/internal/models"
)

var _ KaspiQRBase = (*mtlsClient.MtlsKaspiClient)(nil)

// Common methods for all clients
type KaspiQRBase interface {
	Ping(ctx context.Context) error
	GetTradePoints(ctx context.Context) ([]models.TradePoint, error)
	DeviceRegister(ctx context.Context, deviceID string, tradePointID int64) (models.DeviceToken, error)
	DeviceDelete(ctx context.Context, deviceToken string) error
	CreateQR(ctx context.Context, deviceToken string, amount float64, externalID string) (models.QrToken, error)
	CreateLink(ctx context.Context, deviceToken string, amount float64, externalID string) (models.PaymentData, error)
	GetPaymentStatus(ctx context.Context, qrPaymentId int64) (models.PaymentStatus, error)
}

// IP-based client methods
type KaspiQRIPBased interface {
	KaspiQRBase
	GetClientInfo(ctx context.Context, phoneNumber, deviceToken string) (models.ClientInfo, error)
	CreateRemotePayment(ctx context.Context, amount float64, phoneNumber, deviceToken, comment string) (models.RemotePayment, error)
	CancelRemotePayment(ctx context.Context, qrPaymentID int64, deviceToken string) (models.PaymentStatus, error)
	PaymentReturn(ctx context.Context, deviceToken string, qrPaymentID int64, amount float64) (models.ReturnOperationId, error)
	GetPaymentDetails(ctx context.Context, qrPaymentID int64, deviceToken string) (models.PaymentDetails, error)
}

// mTLS client methods
type KaspiQRMTLS interface {
	KaspiQRBase
	CreateReturn(ctx context.Context, deviceToken, externalID string) (models.Return, error)
	GetReturnStatus(ctx context.Context, qrReturnID int64) (models.ReturnStatus, error)
	ReturnOperations(ctx context.Context, deviceToken string, qrReturnID int64, maxResult int64) ([]models.RecentOperation, error)
	PaymentReturn(ctx context.Context, deviceToken string, qrPaymentID int64, qrReturnID int64, amount float64) (models.ReturnOperationId, error)
	GetPaymentDetails(ctx context.Context, qrPaymentID int64, deviceToken string) (models.PaymentDetails, error)
}
