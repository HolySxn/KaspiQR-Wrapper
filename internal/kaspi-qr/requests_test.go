package kaspihandler

import (
	"context"
	"testing"

	"github.com/HolySxn/KaspiQR-Wrapper/config"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const testAPIKey = "test_F3AD8556847B4736B391CB4D5CFDD14D"
const baseURL = "https://mtokentest.kaspi.kz:8543/r1/v01"

func newTestHandler() (*KaspiHandler, error) {
	cfg := &config.Config{
		Kaspi: struct {
			BaseURL    string          `mapstructure:"base_url"`
			AuthMode   config.AuthMode `mapstructure:"auth_mode"`
			APIKey     string          `mapstructure:"-"`
			ClientCert string          `mapstructure:"-"`
			ClientKey  string          `mapstructure:"-"`
			CACert     string          `mapstructure:"-"`
		}{
			BaseURL:    baseURL,
			AuthMode:   config.AuthModeAPIKey,
			APIKey:     testAPIKey,
			ClientCert: "",
			ClientKey:  "",
			CACert:     "",
		},
	}
	return NewKaspiHandler(cfg)
}

func TestTradePoints(t *testing.T) {
	handler, err := newTestHandler()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	tradePoints, err := handler.GetTradePoints(ctx)
	t.Log(tradePoints)

	assert.NoError(t, err)
	assert.NotNil(t, tradePoints)
}

func TestDeviceRegister(t *testing.T) {
	handler, err := newTestHandler()
	assert.NoError(t, err)

	deviceID := "GFC-456398"
	var tradePoint int64 = 23
	ctx := context.Background()
	deviceToken, err := handler.DeviceRegister(ctx, deviceID, tradePoint)

	assert.NoError(t, err)
	assert.NotEqual(t, "", deviceToken.Token)
}

func TestDeviceDelete(t *testing.T) {
	handler, err := newTestHandler()
	assert.NoError(t, err)

	deviceID := "GFC-456398"
	var tradePoint int64 = 23
	ctx := context.Background()
	deviceToken, err := handler.DeviceRegister(ctx, deviceID, tradePoint)
	assert.NoError(t, err)

	err = handler.DeviceDelete(ctx, deviceToken.Token)
	assert.NoError(t, err)
}

func TestQRCreate(t *testing.T) {
	handler, err := newTestHandler()
	assert.NoError(t, err)

	deviceID := "GFC-456398"
	var tradePoint int64 = 23
	ctx := context.Background()
	deviceToken, err := handler.DeviceRegister(ctx, deviceID, tradePoint)
	assert.NoError(t, err)

	qr, err := handler.CreateQR(ctx, deviceToken.Token, 200, uuid.NewString())
	assert.NoError(t, err)
	assert.NotEmpty(t, qr.Token)
}

func TestLinkCreate(t *testing.T) {
	handler, err := newTestHandler()
	assert.NoError(t, err)

	deviceID := "GFC-456398"
	var tradePoint int64 = 23
	ctx := context.Background()
	deviceToken, err := handler.DeviceRegister(ctx, deviceID, tradePoint)
	assert.NoError(t, err)

	linkData, err := handler.CreateLink(ctx, deviceToken.Token, 200, uuid.NewString())
	assert.NoError(t, err)
	assert.NotEmpty(t, linkData.PaymentLink)
	t.Log(linkData)
}

func TestPaymentStatus(t *testing.T) {
	handler, err := newTestHandler()
	assert.NoError(t, err)

	deviceID := "GFC-456398"
	var tradePoint int64 = 23
	ctx := context.Background()
	deviceToken, err := handler.DeviceRegister(ctx, deviceID, tradePoint)
	assert.NoError(t, err)

	qr, err := handler.CreateQR(ctx, deviceToken.Token, 200, uuid.NewString())
	assert.NoError(t, err)

	status, err := handler.GetPaymentStatus(ctx, deviceToken.Token, qr.Token)
	assert.NoError(t, err)
	assert.NotEmpty(t, status)
	t.Log(status)
}
