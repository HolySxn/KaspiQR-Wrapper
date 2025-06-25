package httpServer

import (
	"log/slog"
	"net/http"

	httpHandler "github.com/HolySxn/KaspiQR-Wrapper/internal/handler/handlers"
	"github.com/gorilla/mux"
)

func NewServer(
	logger *slog.Logger,
	serverHandler *httpHandler.Handler,
) http.Handler {
	router := mux.NewRouter()
	addRoutes(router, serverHandler)

	return router
}
