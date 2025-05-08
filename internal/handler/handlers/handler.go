package httpHandler

import (
	"context"
	"log/slog"

	kaspiqr "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr"
)

type Handler struct {
	logger      *slog.Logger
	kaspiClient kaspiqr.KaspiQR
}

func NewHandler(logger *slog.Logger, kaspiClient kaspiqr.KaspiQR) *Handler {
	return &Handler{
		logger:      logger,
		kaspiClient: kaspiClient,
	}
}

func (h *Handler) QRPay() {
	blankCtx := context.Background()
	deviceToken := "device-token"
	amount := float64(200)
	externalID := "externalID"
	_, err := h.kaspiClient.CreateQR(blankCtx, deviceToken, amount, externalID)
	if err != nil {
		return
	}
}
