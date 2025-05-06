package kaspihandler

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/HolySxn/KaspiQR-Wrapper/config"
	"github.com/HolySxn/KaspiQR-Wrapper/internal/models"
	"github.com/google/uuid"
)

var kaspiStatusMessages = map[int]string{
	0:         "Успешный статус операции",
	-10000:    "Отсутствует сертификат клиента",
	-1501:     "Устройство с заданным идентификатором не найдено",
	-1502:     "Устройство не активно",
	-1503:     "Устройство уже добавлено в другую торговую точку",
	-1601:     "Покупка не найдена",
	-14000002: "Отсутствуют торговые точки",
	-99000002: "Торговая точка не найдена",
	-99000005: "Сумма возврата не может превышать сумму покупки",
	-99000006: "Ошибка возврата",
	-99000018: "Торговая точка отключена",
	-99000026: "Торговая точка не принимает оплату с QR",
	-99000028: "Указана неверная сумма операции",
	-99000033: "Нет доступных методов оплаты",
	-99000001: "Покупка не найдена",
	-99000003: "Торговая точка покупки не соответствует текущему устройству",
	-99000011: "Невозможно вернуть покупку",
	-99000020: "Частичный возврат невозможен",
	-999:      "Сервис временно недоступен",
}

type KaspiHandler struct {
	Client  *http.Client
	baseURL string
	apiKey  string
}

func NewKaspiHandler(cfg *config.Config) (*KaspiHandler, error) {
	var httpClient *http.Client

	switch cfg.Kaspi.AuthMode {
	case config.AuthModeAPIKey:
		httpClient = &http.Client{}
	case config.AuthModeMTLS, config.AuthModeIPBased:
		tlsConfig, err := tlsConfig(cfg.Kaspi.ClientCert, cfg.Kaspi.ClientKey, cfg.Kaspi.CACert)
		if err != nil {
			return nil, err
		}

		httpClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		}

	}

	return &KaspiHandler{
		Client:  httpClient,
		baseURL: cfg.Kaspi.BaseURL,
		apiKey:  cfg.Kaspi.APIKey,
	}, nil
}

func (h *KaspiHandler) seteHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Request-ID", uuid.New().String())
	if h.apiKey != "" {
		req.Header.Set("Api-Key", h.apiKey)
	}
}

func (h *KaspiHandler) doRequest(ctx context.Context, method, url string, body interface{}) (interface{}, error) {
	var reqBody *bytes.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(bodyBytes)
	} else {
		reqBody = bytes.NewReader(nil)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	h.seteHeader(req)

	res, err := h.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	var rawResp models.RawResponse
	if err := json.NewDecoder(res.Body).Decode(&rawResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if rawResp.StatusCode != 0 {
		return nil, fmt.Errorf("%s", kaspiStatusMessages[rawResp.StatusCode])
	}

	return rawResp.Data, nil
}

func tlsConfig(certFile, keyFile, caFile string) (*tls.Config, error) {
	clientCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	caCert, err := os.ReadFile(caFile)
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to append CA cert")
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      caCertPool,
		MinVersion:   tls.VersionTLS12,
	}

	return config, nil
}
