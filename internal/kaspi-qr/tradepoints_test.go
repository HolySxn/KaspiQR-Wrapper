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
	assert.NoError(t, err)
	assert.NotNil(t, tradePoints)
}
