package models

import "time"

type RawResponse struct {
	StatusCode int         `json:"StatusCode"`
	Message    string      `json:"Message"`
	Data       interface{} `json:"Data,omitempty"`
}

type TradePoint struct {
	TradePointID   int64  `json:"TradePointId"`
	TradePointName string `json:"TradePointName"`
}

type DeviceToken struct {
	Token string `json:"DeviceToken"`
}

type QrPaymentBehaviorOptions struct {
	StatusPollingInterval      int `json:"StatusPollingInterval"`
	QrCodeScanWaitTimeout      int `json:"QrCodeScanWaitTimeout"`
	PaymentConfirmationTimeout int `json:"PaymentConfirmationTimeout"`
}

type QrToken struct {
	Token                    string                   `json:"QrToken"`
	ExpireDate               time.Time                `json:"ExpireDate"`
	QrPaymentId              int64                      `json:"QrPaymentId"`
	PaymentMethods           []string                 `json:"PaymentMethods"`
	QrPaymentBehaviorOptions QrPaymentBehaviorOptions `json:"QrPaymentBehaviorOptions"`
}
