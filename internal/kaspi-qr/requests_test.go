package kaspihandler

import (
	"context"
	"crypto/tls"
	"net/http"
	"testing"
	"time"

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
	tradePoints, err := handler.HandleGetTradePoints(ctx)
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
	deviceToken, err := handler.HandleDeviceRegister(ctx, deviceID, tradePoint)

	assert.NoError(t, err)
	assert.Equal(t, "7ba61c32-a110-4ace-ae74-a2f4a81fe6e6", deviceToken.Token)
}
