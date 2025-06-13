package models

type RawResponse struct {
	StatusCode int         `json:"StatusCode"`
	Message    string      `json:"Message"`
	Data       interface{} `json:"Data"`
}

type TradePoint struct {
	TradePointID   int64  `json:"TradePointId"`
	TradePointName string `json:"TradePointName"`
}
