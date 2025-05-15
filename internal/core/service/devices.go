package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/HolySxn/KaspiQR-Wrapper/internal/adapters/repository"
	kaspiqr "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr"
)

type DeviceRepository interface {
	AddDevice(ctx context.Context, deviceID, token, tradePointID string) error
	GetDeviceByID(ctx context.Context, id string) (*repository.Device, error)
	GetDeviceByDeviceID(ctx context.Context, deviceID string) (*repository.Device, error)
	GetDeviceByToken(ctx context.Context, token string) (*repository.Device, error)
	UpdateDeviceToken(ctx context.Context, deviceID, newToken string) error
	UpdateDeviceStatus(ctx context.Context, deviceID, status string) error
	DeleteDevice(ctx context.Context, deviceID string) error
	ListDevices(ctx context.Context, status string) ([]repository.Device, error)
}

type DeviceService struct {
	repo        DeviceRepository
	kaspiClient kaspiqr.KaspiQRBase
	logger      *slog.Logger
}

func NewDeviceService(repo DeviceRepository, kaspiClient kaspiqr.KaspiQRBase, logger *slog.Logger) *DeviceService {
	return &DeviceService{
		repo:        repo,
		kaspiClient: kaspiClient,
		logger:      logger,
	}
}

func (s *DeviceService) RegisterDevice(ctx context.Context, deviceID string, tradePointID int64) (string, error) {
	existingDevice, err := s.repo.GetDeviceByDeviceID(ctx, deviceID)
	if err == nil && existingDevice != nil {
		if existingDevice.Status == "active" {
			s.logger.Info("device already registered and active", slog.String("device_id", deviceID))
			return existingDevice.Token, nil
		} else if existingDevice.Status == "inactive" {
			s.logger.Info("re-registering inactive device", slog.String("device_id", deviceID))
		}
	}

	s.logger.Info("registering new device with Kaspi QR",
		slog.String("device_id", deviceID),
		slog.Int64("trade_point_id", tradePointID),
	)

	deviceToken, err := s.kaspiClient.DeviceRegister(ctx, deviceID, tradePointID)
	if err != nil {
		s.logger.Error("failed to register device with Kaspi QR",
			slog.String("device_id", deviceID),
			slog.String("error", err.Error()),
		)
		return "", fmt.Errorf("failed to register device with Kaspi QR: %w", err)
	}

	tradePointIDStr := fmt.Sprintf("%d", tradePointID)

	if existingDevice != nil {
		err = s.repo.UpdateDeviceToken(ctx, deviceID, deviceToken.Token)
		if err != nil {
			s.logger.Error("failed to update device token in database",
				slog.String("device_id", deviceID),
				slog.String("error", err.Error()),
			)
			return "", fmt.Errorf("failed to update device token in database: %w", err)
		}

		err = s.repo.UpdateDeviceStatus(ctx, deviceID, "active")
		if err != nil {
			s.logger.Error("failed to update device status in database",
				slog.String("device_id", deviceID),
				slog.String("error", err.Error()),
			)
		}
	} else {
		err = s.repo.AddDevice(ctx, deviceID, deviceToken.Token, tradePointIDStr)
		if err != nil {
			s.logger.Error("failed to store device in database",
				slog.String("device_id", deviceID),
				slog.String("error", err.Error()),
			)
			return "", fmt.Errorf("failed to store device in database: %w", err)
		}
	}

	s.logger.Info("device registered successfully",
		slog.String("device_id", deviceID),
		slog.String("trade_point_id", tradePointIDStr),
	)

	return deviceToken.Token, nil
}

func (s *DeviceService) DeleteDevice(ctx context.Context, deviceID string) error {
	device, err := s.repo.GetDeviceByDeviceID(ctx, deviceID)
	if err != nil {
		s.logger.Error("failed to get device from database",
			slog.String("device_id", deviceID),
			slog.String("error", err.Error()),
		)
		return fmt.Errorf("failed to get device from database: %w", err)
	}

	err = s.kaspiClient.DeviceDelete(ctx, device.Token)
	if err != nil {
		s.logger.Error("failed to delete device from Kaspi QR",
			slog.String("device_id", deviceID),
			slog.String("error", err.Error()),
		)
		_ = s.repo.UpdateDeviceStatus(ctx, deviceID, "inactive")

		return fmt.Errorf("failed to delete device from Kaspi QR: %w", err)
	}

	err = s.repo.UpdateDeviceStatus(ctx, deviceID, "inactive")
	if err != nil {
		s.logger.Error("failed to update device status in database",
			slog.String("device_id", deviceID),
			slog.String("error", err.Error()),
		)
		return fmt.Errorf("failed to update device status in database: %w", err)
	}

	s.logger.Info("device deleted successfully", slog.String("device_id", deviceID))
	return nil
}

func (s *DeviceService) GetDeviceToken(ctx context.Context, deviceID string) (string, error) {
	device, err := s.repo.GetDeviceByDeviceID(ctx, deviceID)
	if err != nil {
		s.logger.Error("failed to get device from database",
			slog.String("device_id", deviceID),
			slog.String("error", err.Error()),
		)
		return "", fmt.Errorf("failed to get device from database: %w", err)
	}

	if device.Status != "active" {
		s.logger.Warn("inactive device token requested", slog.String("device_id", deviceID))
		return "", fmt.Errorf("device is not active")
	}

	return device.Token, nil
}

func (s *DeviceService) ListDevices(ctx context.Context, status string) ([]repository.Device, error) {
	devices, err := s.repo.ListDevices(ctx, status)
	if err != nil {
		s.logger.Error("failed to list devices from database",
			slog.String("status_filter", status),
			slog.String("error", err.Error()),
		)
		return nil, fmt.Errorf("failed to list devices: %w", err)
	}

	return devices, nil
}
