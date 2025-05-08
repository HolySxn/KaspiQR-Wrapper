package models

import "time"

type RawResponse struct {
	StatusCode int         `json:"StatusCode"`
	Message    string      `json:"Message"`
	Data       interface{} `json:"Data,omitempty"`
}

type TradePoint struct {
	TradePointId   int64  `json:"TradePointId"`
	TradePointName string `json:"TradePointName"`
}

type DeviceToken struct {
	Token string `json:"DeviceToken"`
}

// Create QR response
type QrPaymentBehaviorOptions struct {
	StatusPollingInterval      int `json:"StatusPollingInterval"`
	QrCodeScanWaitTimeout      int `json:"QrCodeScanWaitTimeout"`
	PaymentConfirmationTimeout int `json:"PaymentConfirmationTimeout"`
}

type QrToken struct {
	Token                    string                   `json:"QrToken"`
	ExpireDate               time.Time                `json:"ExpireDate"`
	QrPaymentId              int64                    `json:"QrPaymentId"`
	PaymentMethods           []string                 `json:"PaymentMethods"`
	QrPaymentBehaviorOptions QrPaymentBehaviorOptions `json:"QrPaymentBehaviorOptions"`
}

// Create link response
type PaymentBehaviorOptions struct {
	StatusPollingInterval      int `json:"StatusPollingInterval"`
	LinkActivationWaitTimeout  int `json:"LinkActivationWaitTimeout"`
	PaymentConfirmationTimeout int `json:"PaymentConfirmationTimeout"`
}

type PaymentData struct {
	PaymentLink            string                 `json:"PaymentLink"`
	ExpireDate             time.Time              `json:"ExpireDate"`
	PaymentId              int                    `json:"PaymentId"`
	PaymentMethods         []string               `json:"PaymentMethods"`
	PaymentBehaviorOptions PaymentBehaviorOptions `json:"PaymentBehaviorOptions"`
}

// QR/Payment status
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

// Create return response
type QrReturnBehaviorOptions struct {
	QrCodeScanEventPollingInterval int `json:"QrCodeScanEventPollingInterval"`
	QrCodeScanWaitTimeout          int `json:"QrCodeScanWaitTimeout"`
}

type Return struct {
	QrToken                 string                  `json:"QrToken"`
	ExpireDate              time.Time               `json:"ExpireDate"`
	QrReturnId              int                     `json:"QrReturnId"`
	QrReturnBehaviorOptions QrReturnBehaviorOptions `json:"QrReturnBehaviorOptions"`
}

// Return status
type ReturnStatus struct {
	Status string `json:"Status"`
}

// Recent purchases response
type RecentOperation struct {
	QrPaymentId     int64     `json:"QrPaymentId"`
	TransactionDate time.Time `json:"TransactionDate"`
	Amount          float64   `json:"Amount"`
}

// Payment details response
type PaymentDetails struct {
	QrPaymentId           int64     `json:"QrPaymentId"`
	TotalAmount           float64   `json:"TotalAmount"`
	AvailableReturnAmount float64   `json:"AvailableReturnAmount"`
	TransactionDate       time.Time `json:"TransactionDate"`
}

// Payment Return response
type ReturnOperationId struct {
	Id int64 `json:"ReturnOperationId"`
}

type ClientInfo struct {
	ClientName string `json:"ClientName"`
}

type RemotePayment struct {
	QrPaymentId int64 `json:"QrPaymentId"`
}

type CancelRemotePaymentStatus struct {
	Status string `json:"Status"`
}
