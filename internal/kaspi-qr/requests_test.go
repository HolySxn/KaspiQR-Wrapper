package kaspihandler

import (
	"context"
	"crypto/tls"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const test_api = "test_F3AD8556847B4736B391CB4D5CFDD14D"
const base_url = "https://mtokentest.kaspi.kz:8543/r1/v01"

func TestTradePoints(t *testing.T) {
	handler := NewKaspiHandler(base_url, test_api)
	handler.Client = &http.Client{
		Timeout: time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	ctx := context.Background()
	tradePoints, err := handler.GetTradePoints(ctx)
	t.Log(tradePoints)

	assert.NoError(t, err)
	assert.NotNil(t, tradePoints)
}

func TestDeviceRegister(t *testing.T) {
	handler := NewKaspiHandler(base_url, test_api)
	handler.Client = &http.Client{
		Timeout: time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	deviceID := "GFC-456398"
	var tradePoint int64 = 23
	ctx := context.Background()
	deviceToken, err := handler.DeviceRegister(ctx, deviceID, tradePoint)

	assert.NoError(t, err)
	assert.NotEqual(t, "", deviceToken)
}

func TestDeviceDelete(t *testing.T) {
	handler := NewKaspiHandler(base_url, test_api)
	handler.Client = &http.Client{
		Timeout: time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	deviceID := "GFC-456398"
	var tradePoint int64 = 23
	ctx := context.Background()
	deviceToken, err := handler.DeviceRegister(ctx, deviceID, tradePoint)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.DeviceDelete(ctx, deviceToken.Token)

	assert.NoError(t, err)
}

func TestQRCreate(t *testing.T) {
	handler := NewKaspiHandler(base_url, test_api)
	handler.Client = &http.Client{
		Timeout: time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	deviceID := "GFC-456398"
	var tradePoint int64 = 23
	ctx := context.Background()
	deviceToken, err := handler.DeviceRegister(ctx, deviceID, tradePoint)
	if err != nil {
		t.Fatal(err)
	}

	qr, err := handler.CreateQR(ctx, deviceToken.Token, 200, uuid.NewString())
	assert.NoError(t, err)
	assert.NotEmpty(t, qr)
}
