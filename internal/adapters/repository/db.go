package repository

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type Queries struct {
	db     DBTX
	logger *slog.Logger
}

func New(db DBTX, logger *slog.Logger) *Queries {
	return &Queries{
		db:     db,
		logger: logger,
	}
}

func WithTx(tx pgx.Tx, logger *slog.Logger) *Queries {
	return &Queries{
		db:     tx,
		logger: logger,
	}

}
