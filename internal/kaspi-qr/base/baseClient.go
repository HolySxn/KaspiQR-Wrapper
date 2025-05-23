package base

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/HolySxn/KaspiQR-Wrapper/internal/models"
	"github.com/google/uuid"
)

var KaspiStatusMessages = map[int]string{
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

type BaseKaspiClient struct {
	Client     *http.Client
	BaseURL    string
	HeadSetter func(req *http.Request)
}

func NewBaseKaspiClient(baseURL string, client *http.Client) *BaseKaspiClient {
	base := &BaseKaspiClient{
		BaseURL: baseURL,
		Client:  client,
	}
	base.HeadSetter = base.setHeader
	return base
}

func (c *BaseKaspiClient) setHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Request-ID", uuid.New().String())
}

func (c *BaseKaspiClient) DoRequest(ctx context.Context, method, url string, body interface{}) (interface{}, error) {
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

	c.HeadSetter(req)

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()

	bodyBytes, _ := io.ReadAll(res.Body)

	// if res.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("Kaspi API error: %s", string(bodyBytes))
	// }

	var rawResp models.RawResponse
	if err := json.Unmarshal(bodyBytes, &rawResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w\nRaw body: %s", err, string(bodyBytes))
	}

	if rawResp.StatusCode != 0 {
		return nil, fmt.Errorf("%s", KaspiStatusMessages[rawResp.StatusCode])
	}

	return rawResp.Data, nil
}
