package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/HolySxn/KaspiQR-Wrapper/config"
	httpServer "github.com/HolySxn/KaspiQR-Wrapper/internal/handler"
	httpHandler "github.com/HolySxn/KaspiQR-Wrapper/internal/handler/handlers"
	kaspihandler "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := config.NewLogger()
	ctx := context.Background()
	cfg, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		logger.Error("failed to load config", slog.Any("error", err))
		return
	}

	kaspiHandler := kaspihandler.NewKaspiHandler(cfg.Kaspi.BaseURL, cfg.Kaspi.APIKey)
	serverHandler := httpHandler.NewHandler(logger, kaspiHandler)

	srv := httpServer.NewServer(logger, serverHandler)
	run(ctx, cfg, srv, logger)
}

func run(ctx context.Context, cfg *config.Config, srv *gin.Engine, logger *slog.Logger) {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(cfg.Server.Host, cfg.Server.Port),
		Handler: srv,
	}

	go func() {
		slog.Info("listening", "addr", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Info("error listening and serving", "error", err)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		logger.Info("Gracefully shutting down...")
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			slog.Info("error shutting down http server", "error", err)
		}
	}()
	wg.Wait()
}
