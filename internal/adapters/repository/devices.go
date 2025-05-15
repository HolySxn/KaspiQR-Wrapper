package repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
)

// Device represents a device record in the database
type Device struct {
	ID           string    `json:"id"`
	DeviceID     string    `json:"device_id"`
	Token        string    `json:"token"`
	TradePointID string    `json:"trade_point_id"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

func (q *Queries) AddDevice(ctx context.Context, deviceID, token, tradePointID string) error {
	query := `INSERT INTO devices (device_id, token, trade_point_id, status)
	VALUES ($1, $2, $3, $4)`
	_, err := q.db.Exec(ctx, query, deviceID, token, tradePointID, "active")
	if err != nil {
		q.logger.Error("failed to add device", slog.String("device_id", deviceID), slog.String("error", err.Error()))
		return fmt.Errorf("failed to add device: %w", err)
	}
	return nil
}

func (q *Queries) GetDeviceByID(ctx context.Context, id string) (*Device, error) {
	query := `SELECT id, device_id, token, trade_point_id, status, created_at FROM devices WHERE id = $1`
	var device Device
	err := q.db.QueryRow(ctx, query, id).Scan(
		&device.ID,
		&device.DeviceID,
		&device.Token,
		&device.TradePointID,
		&device.Status,
		&device.CreatedAt,
	)
	if err != nil {
		q.logger.Error("failed to get device by ID", slog.String("id", id), slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to get device by ID: %w", err)
	}
	return &device, nil
}

func (q *Queries) GetDeviceByDeviceID(ctx context.Context, deviceID string) (*Device, error) {
	query := `SELECT id, device_id, token, trade_point_id, status, created_at FROM devices WHERE device_id = $1`
	var device Device
	err := q.db.QueryRow(ctx, query, deviceID).Scan(
		&device.ID,
		&device.DeviceID,
		&device.Token,
		&device.TradePointID,
		&device.Status,
		&device.CreatedAt,
	)
	if err != nil {
		q.logger.Error("failed to get device by device_id", slog.String("device_id", deviceID), slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to get device by device_id: %w", err)
	}
	return &device, nil
}

func (q *Queries) GetDeviceByToken(ctx context.Context, token string) (*Device, error) {
	query := `SELECT id, device_id, token, trade_point_id, status, created_at FROM devices WHERE token = $1`
	var device Device
	err := q.db.QueryRow(ctx, query, token).Scan(
		&device.ID,
		&device.DeviceID,
		&device.Token,
		&device.TradePointID,
		&device.Status,
		&device.CreatedAt,
	)
	if err != nil {
		q.logger.Error("failed to get device by token", slog.String("token", token), slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to get device by token: %w", err)
	}
	return &device, nil
}

func (q *Queries) UpdateDeviceToken(ctx context.Context, deviceID, newToken string) error {
	query := `UPDATE devices SET token = $1 WHERE device_id = $2`
	result, err := q.db.Exec(ctx, query, newToken, deviceID)
	if err != nil {
		q.logger.Error("failed to update device token", slog.String("device_id", deviceID), slog.String("error", err.Error()))
		return fmt.Errorf("failed to update device token: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("no device found with device_id: %s", deviceID)
	}
	return nil
}

func (q *Queries) UpdateDeviceStatus(ctx context.Context, deviceID, status string) error {
	query := `UPDATE devices SET status = $1 WHERE device_id = $2`
	result, err := q.db.Exec(ctx, query, status, deviceID)
	if err != nil {
		q.logger.Error("failed to update device status", slog.String("device_id", deviceID), slog.String("status", status), slog.String("error", err.Error()))
		return fmt.Errorf("failed to update device status: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("no device found with device_id: %s", deviceID)
	}
	return nil
}

func (q *Queries) DeleteDevice(ctx context.Context, deviceID string) error {
	query := `DELETE FROM devices WHERE device_id = $1`
	result, err := q.db.Exec(ctx, query, deviceID)
	if err != nil {
		q.logger.Error("failed to delete device", slog.String("device_id", deviceID), slog.String("error", err.Error()))
		return fmt.Errorf("failed to delete device: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("no device found with device_id: %s", deviceID)
	}
	return nil
}

func (q *Queries) ListDevices(ctx context.Context, status string) ([]Device, error) {
	var (
		query string
		rows  pgx.Rows
		err   error
	)
	if status != "" {
		query = `SELECT id, device_id, token, trade_point_id, status, created_at FROM devices WHERE status = $1 ORDER BY created_at DESC`
		rows, err = q.db.Query(ctx, query, status)
	} else {
		query = `SELECT id, device_id, token, trade_point_id, status, created_at FROM devices ORDER BY created_at DESC`
		rows, err = q.db.Query(ctx, query)
	}
	if err != nil {
		q.logger.Error("failed to list devices", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to list devices: %w", err)
	}
	defer rows.Close()
	var devices []Device
	for rows.Next() {
		var device Device
		err := rows.Scan(
			&device.ID,
			&device.DeviceID,
			&device.Token,
			&device.TradePointID,
			&device.Status,
			&device.CreatedAt,
		)
		if err != nil {
			q.logger.Error("failed to scan device row", slog.String("error", err.Error()))
			return nil, fmt.Errorf("failed to scan device row: %w", err)
		}
		devices = append(devices, device)
	}
	if err := rows.Err(); err != nil {
		q.logger.Error("error iterating through device rows", slog.String("error", err.Error()))
		return nil, fmt.Errorf("error iterating through device rows: %w", err)
	}
	return devices, nil
}
