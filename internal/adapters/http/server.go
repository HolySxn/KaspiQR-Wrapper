package httpServer

import (
	"log/slog"
	"net/http"

	httpHandler "github.com/HolySxn/KaspiQR-Wrapper/internal/adapters/http/handlers"
	"github.com/gorilla/mux"
)

func NewServer(
	logger *slog.Logger,
	serverHandler *httpHandler.Handler,
	authMode string,
) http.Handler {
	router := mux.NewRouter()
	addRoutes(router, serverHandler, authMode)

	return router
}
