package httpServer

import (
	"github.com/HolySxn/KaspiQR-Wrapper/config"
	httpHandler "github.com/HolySxn/KaspiQR-Wrapper/internal/handler/handlers"
	"github.com/gorilla/mux"
)

func addRoutes(router *mux.Router, serverHandler *httpHandler.Handler, authMode string) {
	router.HandleFunc("/health/kaspiapi", serverHandler.Ping).Methods("GET")
	switch authMode {
	case config.AuthModeAPIKey:
		r1 := router.PathPrefix("/r1").Subrouter()
		r1.HandleFunc("/partner/tradepoints", serverHandler.GetTradePoints).Methods("GET")
		r1.HandleFunc("/device/register", serverHandler.DeviceRegister).Methods("POST")
		r1.HandleFunc("/device/delete", serverHandler.DeviceDelete).Methods("POST")
		r1.HandleFunc("/qr/create", serverHandler.CreateQR).Methods("POST")
		r1.HandleFunc("/qr/create-link", serverHandler.CreateLink).Methods("POST")
		r1.HandleFunc("/payment/status/{QrPaymentId}", serverHandler.GetPaymentStatus).Methods("GET")
	case config.AuthModeMTLS:
		r2 := router.PathPrefix("/r2").Subrouter()
		r2.HandleFunc("/partner/tradepoints", serverHandler.GetTradePoints).Methods("GET")
		r2.HandleFunc("/device/register", serverHandler.DeviceRegister).Methods("POST")
		r2.HandleFunc("/device/delete", serverHandler.DeviceDelete).Methods("POST")
		r2.HandleFunc("/qr/create", serverHandler.CreateQR).Methods("POST")
		r2.HandleFunc("/qr/create-link", serverHandler.CreateLink).Methods("POST")
		r2.HandleFunc("/payment/status/{QrPaymentId}", serverHandler.GetPaymentStatus).Methods("GET")
		r2.HandleFunc("/return/create", serverHandler.CreateReturnMTLS).Methods("POST")
		r2.HandleFunc("/return/status/{QrReturnId}", serverHandler.GetReturnStatusMTLS).Methods("GET")
		r2.HandleFunc("/return/operations", serverHandler.ReturnOperationsMTLS).Methods("POST")
		r2.HandleFunc("/payment/details", serverHandler.GetPaymentDetailsMTLS).Methods("GET")
		r2.HandleFunc("/payment/return", serverHandler.PaymentReturnMTLS).Methods("POST")
	case config.AuthModeIPBased:
		r3 := router.PathPrefix("/r3").Subrouter()
		r3.HandleFunc("/partner/tradepoints", serverHandler.GetTradePoints).Methods("GET")
		r3.HandleFunc("/device/register", serverHandler.DeviceRegister).Methods("POST")
		r3.HandleFunc("/device/delete", serverHandler.DeviceDelete).Methods("POST")
		r3.HandleFunc("/qr/create", serverHandler.CreateQR).Methods("POST")
		r3.HandleFunc("/qr/create-link", serverHandler.CreateLink).Methods("POST")
		r3.HandleFunc("/payment/status/{QrPaymentId}", serverHandler.GetPaymentStatus).Methods("GET")
		r3.HandleFunc("/payment/details", serverHandler.GetPaymentDetailsIPBased).Methods("GET")
		r3.HandleFunc("/payment/return", serverHandler.PaymentReturnIPBased).Methods("POST")
		r3.HandleFunc("/remote/client-info", serverHandler.GetClientInfoIPBased).Methods("GET")
		r3.HandleFunc("/remote/create", serverHandler.CreateRemotePaymentIPBased).Methods("POST")
		r3.HandleFunc("/remote/cancel", serverHandler.CancelRemotePaymentIPBased).Methods("POST")
	}
}
