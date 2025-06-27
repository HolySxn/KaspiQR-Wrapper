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
	httpServer "github.com/HolySxn/KaspiQR-Wrapper/internal/adapters/http"
	httpHandler "github.com/HolySxn/KaspiQR-Wrapper/internal/adapters/http/handlers"
	"github.com/HolySxn/KaspiQR-Wrapper/internal/adapters/repository"
	"github.com/HolySxn/KaspiQR-Wrapper/internal/core/service"
	kaspiqr "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr"
)

func main() {
	initialLogger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(initialLogger)

	slog.Info("starting service")

	ctx := context.Background()
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		slog.Error("failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	deps, err := config.NewDependencies(
		ctx,
		config.WithLogger(cfg.Server.LogLvl),
		config.WithPosgres(
			cfg.Postgres.User,
			cfg.Postgres.Password,
			cfg.Postgres.Host,
			cfg.Postgres.Port,
			cfg.Postgres.DBName,
			cfg.Postgres.SSLMode,
		))

	if err != nil {
		slog.Error("failed to load dependencies", slog.Any("error", err))
		os.Exit(1)
	}
	defer deps.Close()
	slog.SetDefault(deps.Logger)

	kaspiClient, err := kaspiqr.NewKaspiClient(cfg)
	if err != nil {
		slog.Error("failed to create kaspi handler", slog.Any("err", err))
		os.Exit(1)
	}

	repo := repository.New(deps.Postgres, deps.Logger)
	deviceService := service.NewDeviceService(repo, kaspiClient, deps.Logger)

	serverHandler := httpHandler.NewHandler(deps.Logger, kaspiClient, deviceService)

	srv := httpServer.NewServer(deps.Logger, serverHandler, cfg.Kaspi.AuthMode)
	run(ctx, cfg, srv)
}

func run(ctx context.Context, cfg *config.Config, srv http.Handler) {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(cfg.Server.Host, cfg.Server.Port),
		Handler: srv,
	}

	go func() {
		slog.Info("server listening",
			"address", httpServer.Addr,
			"auth_mode", cfg.Kaspi.AuthMode)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Info("error listening and serving", "error", err)
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
		slog.Info("Gracefully shutting down...")
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			slog.Info("error shutting down http server", "error", err)
		}
	}()
	wg.Wait()
}
