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
	ipClient, ok := h.kaspiClient.(kaspiqr.KaspiQRIPBased)
	if !ok {
		http.Error(w, "Not supported for this client type", http.StatusNotImplemented)
		return
	}

	phone := r.URL.Query().Get("phoneNumber")
	deviceToken := r.URL.Query().Get("deviceToken")

	ctx := r.Context()

	info, err := ipClient.GetClientInfo(ctx, phone, deviceToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(info)
}

// POST /remote/create
func (h *Handler) CreateRemotePaymentIPBased(w http.ResponseWriter, r *http.Request) {
	ipClient, ok := h.kaspiClient.(kaspiqr.KaspiQRIPBased)
	if !ok {
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
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	payment, err := ipClient.CreateRemotePayment(ctx, req.Amount, req.PhoneNumber, req.DeviceToken, req.Comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(payment)
}

// POST /remote/cancel
func (h *Handler) CancelRemotePaymentIPBased(w http.ResponseWriter, r *http.Request) {
	ipClient, ok := h.kaspiClient.(kaspiqr.KaspiQRIPBased)
	if !ok {
		http.Error(w, "Not supported for this client type", http.StatusNotImplemented)
		return
	}

	var req struct {
		QrPaymentID int64  `json:"qrPaymentId"`
		DeviceToken string `json:"deviceToken"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	status, err := ipClient.CancelRemotePayment(ctx, req.QrPaymentID, req.DeviceToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(status)
}

// POST /payment/return
func (h *Handler) PaymentReturnIPBased(w http.ResponseWriter, r *http.Request) {
	ipClient, ok := h.kaspiClient.(kaspiqr.KaspiQRIPBased)
	if !ok {
		http.Error(w, "Not supported for this client type", http.StatusNotImplemented)
		return
	}

	var req struct {
		DeviceToken string  `json:"deviceToken"`
		QrPaymentID int64   `json:"qrPaymentId"`
		Amount      float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	result, err := ipClient.PaymentReturn(ctx, req.DeviceToken, req.QrPaymentID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// GET /payment/details
func (h *Handler) GetPaymentDetailsIPBased(w http.ResponseWriter, r *http.Request) {
	ipClient, ok := h.kaspiClient.(kaspiqr.KaspiQRIPBased)
	if !ok {
		http.Error(w, "Not supported for this client type", http.StatusNotImplemented)
		return
	}

	qrPaymentIDStr := r.URL.Query().Get("QrPaymentId")
	deviceToken := r.URL.Query().Get("DeviceToken")

	qrPaymentID, err := strconv.ParseInt(qrPaymentIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid qrPaymentId", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	details, err := ipClient.GetPaymentDetails(ctx, qrPaymentID, deviceToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(details)
}
