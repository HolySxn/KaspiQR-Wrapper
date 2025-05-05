package httpHandler

import (
	"log/slog"

	kaspihandler "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr"
)

type Handler struct {
	logger       *slog.Logger
	kaspiHandler *kaspihandler.KaspiHandler
}

func NewHandler(logger *slog.Logger, kaspiHandler *kaspihandler.KaspiHandler) *Handler {
	return &Handler{
		logger:       logger,
		kaspiHandler: kaspiHandler,
	}
}
