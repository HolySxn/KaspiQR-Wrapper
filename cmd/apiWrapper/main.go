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
	kaspiqr "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr"
)

func main() {
	logger := config.NewLogger()
	ctx := context.Background()
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		logger.Error("failed to load config", slog.Any("error", err))
		return
	}

	kaspiClient, err := kaspiqr.NewKaspiClient(cfg)
	if err != nil {
		logger.Error("failed to create kaspi handler", slog.Any("err", err))
		return
	}
	serverHandler := httpHandler.NewHandler(logger, kaspiClient)

	srv := httpServer.NewServer(logger, serverHandler, cfg.Kaspi.AuthMode)
	run(ctx, cfg, srv, logger)
}

func run(ctx context.Context, cfg *config.Config, srv http.Handler, logger *slog.Logger) {
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
