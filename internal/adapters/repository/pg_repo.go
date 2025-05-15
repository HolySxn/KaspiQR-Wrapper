package repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewPostgresRepository(db *pgxpool.Pool, logger *slog.Logger) *PostgresRepository {
	return &PostgresRepository{
		db:     db,
		logger: logger,
	}
}
