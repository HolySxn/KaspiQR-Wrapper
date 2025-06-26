package base

import (
	"context"
	"net/http"
)

func (c *BaseKaspiClient) Ping(ctx context.Context) error {
	// Perform a health check by sending a request to the health endpoint
	url := c.BaseURL + "/health/ping"

	_, err := c.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	return nil
}
