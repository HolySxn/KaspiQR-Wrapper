package config

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Dependencies struct {
	Postgres *pgxpool.Pool
	Logger   *slog.Logger
}

type Option func(context.Context, *Dependencies) error

func (d *Dependencies) Close() {
	if d == nil {
		return
	}

	if d.Postgres != nil {
		d.Postgres.Close()
	}
}

func NewDependencies(ctx context.Context, opts ...Option) (deps *Dependencies, err error) {
	defer func() {
		if err != nil {
			deps.Close()
		}
	}()

	deps = &Dependencies{}

	for _, opt := range opts {
		if err := opt(ctx, deps); err != nil {
			return nil, err
		}
	}

	return deps, nil
}

func WithPosgres(
	user string,
	password string,
	host string,
	port string,
	dbName string,
	sslmode string,
) Option {
	return func(ctx context.Context, d *Dependencies) error {
		format := "postgresql://%s:%s@%s:%s/%s?sslmode=%s"
		connString := fmt.Sprintf(format, user, password, host, port, dbName, sslmode)

		pool, err := pgxpool.New(ctx, connString)
		if err != nil {
			return err
		}
		slog.Info("DB connection established")

		d.Postgres = pool
		return nil
	}
}

const (
	EnvDev  = "dev"
	EnvProd = "prod"
)

func WithLogger(level string) Option {
	return func(ctx context.Context, d *Dependencies) error {
		var logLvl slog.Level

		switch level {
		case EnvDev:
			logLvl = slog.LevelDebug
		case EnvProd:
			logLvl = slog.LevelInfo
		}

		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLvl,
		}))

		d.Logger = logger
		return nil
	}
}
