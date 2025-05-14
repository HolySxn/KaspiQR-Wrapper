package httpHandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	kaspiqr "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr"
)

// --- IP-based handlers ---

// GET /remote/client-info
func (h *Handler) GetClientInfoIPBased(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling GetClientInfoIPBased request", "method", r.Method, "url", r.URL.String())

	ipClient, ok := h.kaspiClient.(kaspiqr.KaspiQRIPBased)
	if !ok {
		h.logger.Error("Client type not supported for GetClientInfoIPBased")
		http.Error(w, "Not supported for this client type", http.StatusNotImplemented)
		return
	}

	phone := r.URL.Query().Get("phoneNumber")
	deviceToken := r.URL.Query().Get("deviceToken")

	ctx := r.Context()

	info, err := ipClient.GetClientInfo(ctx, phone, deviceToken)
	if err != nil {
		h.logger.Error("Failed to retrieve client information", "error", err)
		http.Error(w, "Failed to retrieve client information. Please try again later.", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Successfully retrieved client information", "phone", phone)
	json.NewEncoder(w).Encode(info)
}

// POST /remote/create
func (h *Handler) CreateRemotePaymentIPBased(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling CreateRemotePaymentIPBased request", "method", r.Method, "url", r.URL.String())

	ipClient, ok := h.kaspiClient.(kaspiqr.KaspiQRIPBased)
	if !ok {
		h.logger.Error("Client type not supported for CreateRemotePaymentIPBased")
		http.Error(w, "Not supported for this client type", http.StatusNotImplemented)
		return
	}

	var req struct {
		Amount      float64 `json:"amount"`
		PhoneNumber string  `json:"phoneNumber"`
		DeviceToken string  `json:"deviceToken"`
		Comment     string  `json:"comment"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Invalid request payload for CreateRemotePaymentIPBased", "error", err)
		http.Error(w, "Invalid request. Please check your input and try again.", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	payment, err := ipClient.CreateRemotePayment(ctx, req.Amount, req.PhoneNumber, req.DeviceToken, req.Comment)
	if err != nil {
		h.logger.Error("Failed to create payment", "error", err)
		http.Error(w, "Failed to create payment. Please try again later.", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Successfully created payment", "amount", req.Amount, "phoneNumber", req.PhoneNumber)
	json.NewEncoder(w).Encode(payment)
}

// POST /remote/cancel
func (h *Handler) CancelRemotePaymentIPBased(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling CancelRemotePaymentIPBased request", "method", r.Method, "url", r.URL.String())

	ipClient, ok := h.kaspiClient.(kaspiqr.KaspiQRIPBased)
	if !ok {
		h.logger.Error("Client type not supported for CancelRemotePaymentIPBased")
		http.Error(w, "Not supported for this client type", http.StatusNotImplemented)
		return
	}

	var req struct {
		QrPaymentID int64  `json:"qrPaymentId"`
		DeviceToken string `json:"deviceToken"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Invalid request payload for CancelRemotePaymentIPBased", "error", err)
		http.Error(w, "Invalid request. Please check your input and try again.", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	status, err := ipClient.CancelRemotePayment(ctx, req.QrPaymentID, req.DeviceToken)
	if err != nil {
		h.logger.Error("Failed to cancel payment", "error", err)
		http.Error(w, "Failed to cancel payment. Please try again later.", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Successfully canceled payment", "qrPaymentId", req.QrPaymentID)
	json.NewEncoder(w).Encode(status)
}

// POST /payment/return
func (h *Handler) PaymentReturnIPBased(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling PaymentReturnIPBased request", "method", r.Method, "url", r.URL.String())

	ipClient, ok := h.kaspiClient.(kaspiqr.KaspiQRIPBased)
	if !ok {
		h.logger.Error("Client type not supported for PaymentReturnIPBased")
		http.Error(w, "Not supported for this client type", http.StatusNotImplemented)
		return
	}

	var req struct {
		DeviceToken string  `json:"deviceToken"`
		QrPaymentID int64   `json:"qrPaymentId"`
		Amount      float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Invalid request payload for PaymentReturnIPBased", "error", err)
		http.Error(w, "Invalid request. Please check your input and try again.", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	result, err := ipClient.PaymentReturn(ctx, req.DeviceToken, req.QrPaymentID, req.Amount)
	if err != nil {
		h.logger.Error("Failed to process payment return", "error", err)
		http.Error(w, "Failed to process payment return. Please try again later.", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Successfully processed payment return", "qrPaymentId", req.QrPaymentID, "amount", req.Amount)
	json.NewEncoder(w).Encode(result)
}

// GET /payment/details
func (h *Handler) GetPaymentDetailsIPBased(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling GetPaymentDetailsIPBased request", "method", r.Method, "url", r.URL.String())

	ipClient, ok := h.kaspiClient.(kaspiqr.KaspiQRIPBased)
	if !ok {
		h.logger.Error("Client type not supported for GetPaymentDetailsIPBased")
		http.Error(w, "Not supported for this client type", http.StatusNotImplemented)
		return
	}

	qrPaymentIDStr := r.URL.Query().Get("QrPaymentId")
	deviceToken := r.URL.Query().Get("DeviceToken")

	qrPaymentID, err := strconv.ParseInt(qrPaymentIDStr, 10, 64)
	if err != nil {
		h.logger.Error("Invalid qrPaymentId for GetPaymentDetailsIPBased", "error", err)
		http.Error(w, "Invalid qrPaymentId. Please check your input and try again.", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	details, err := ipClient.GetPaymentDetails(ctx, qrPaymentID, deviceToken)
	if err != nil {
		h.logger.Error("Failed to retrieve payment details", "error", err)
		http.Error(w, "Failed to retrieve payment details. Please try again later.", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Successfully retrieved payment details", "qrPaymentId", qrPaymentID)
	json.NewEncoder(w).Encode(details)
}
