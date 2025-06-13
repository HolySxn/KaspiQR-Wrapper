package httpServer

import (
	"log/slog"

	httpHandler "github.com/HolySxn/KaspiQR-Wrapper/internal/handler/handlers"
	"github.com/gin-gonic/gin"
)

func NewServer(
	logger *slog.Logger,
	serverHandler *httpHandler.Handler,
) *gin.Engine {
	router := gin.New()
	addRoutes(router, serverHandler)

	return router
}
