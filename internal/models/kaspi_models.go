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

type PaymentBehaviorOptions struct {
	StatusPollingInterval      int `json:"StatusPollingInterval"`
	LinkActivationWaitTimeout  int `json:"LinkActivationWaitTimeout"`
	PaymentConfirmationTimeout int `json:"PaymentConfirmationTimeout"`
}

type PaymentData struct {
	PaymentLink             string                `json:"PaymentLink"`
	ExpireDate              time.Time             `json:"ExpireDate"`
	PaymentId               int                   `json:"PaymentId"`
	PaymentMethods          []string              `json:"PaymentMethods"`
	PaymentBehaviorOptions  PaymentBehaviorOptions `json:"PaymentBehaviorOptions"`
}

type PaymentStatus struct {
	Status        string      `json:"Status"`
	TransactionId string      `json:"TransactionId"`
	LoanOfferName string      `json:"LoanOfferName"`
	LoanTerm      int         `json:"LoanTerm"`
	IsOffer       bool        `json:"IsOffer"`
	ProductType   string      `json:"ProductType"`
	Data          PaymentData `json:"Data"`
	Amount        float64     `json:"Amount"`
	StoreName     string      `json:"StoreName"`
	Address       string      `json:"Address"`
	City          string      `json:"City"`
}