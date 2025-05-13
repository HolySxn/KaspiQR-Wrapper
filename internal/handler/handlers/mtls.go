package httpHandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	kaspiqr "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr"
	"github.com/gorilla/mux"
)

// --- mTLS handlers ---

// POST /return/create
func (h *Handler) CreateReturnMTLS(w http.ResponseWriter, r *http.Request) {
	mtlsClient, ok := h.kaspiClient.(kaspiqr.KaspiQRMTLS)
	if !ok {
		http.Error(w, "Not supported for this client type", http.StatusNotImplemented)
		return
	}

	var req struct {
		DeviceToken string `json:"deviceToken"`
		ExternalID  string `json:"externalId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	result, err := mtlsClient.CreateReturn(ctx, req.DeviceToken, req.ExternalID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// GET /return/status/{qrReturnId}
func (h *Handler) GetReturnStatusMTLS(w http.ResponseWriter, r *http.Request) {
	mtlsClient, ok := h.kaspiClient.(kaspiqr.KaspiQRMTLS)
	if !ok {
		http.Error(w, "Not supported for this client type", http.StatusNotImplemented)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["QrReturnId"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid qrReturnId", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	status, err := mtlsClient.GetReturnStatus(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(status)
}

// POST /return/operations
func (h *Handler) ReturnOperationsMTLS(w http.ResponseWriter, r *http.Request) {
	mtlsClient, ok := h.kaspiClient.(kaspiqr.KaspiQRMTLS)
	if !ok {
		http.Error(w, "Not supported for this client type", http.StatusNotImplemented)
		return
	}

	var req struct {
		DeviceToken string `json:"deviceToken"`
		QrReturnID  int64  `json:"qrReturnId"`
		MaxResult   int64  `json:"maxResult"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	ops, err := mtlsClient.ReturnOperations(ctx, req.DeviceToken, req.QrReturnID, req.MaxResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ops)
}

// GET /payment/details
func (h *Handler) GetPaymentDetailsMTLS(w http.ResponseWriter, r *http.Request) {
	mtlsClient, ok := h.kaspiClient.(kaspiqr.KaspiQRMTLS)
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

	details, err := mtlsClient.GetPaymentDetails(ctx, qrPaymentID, deviceToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(details)
}

// POST /payment/return
func (h *Handler) PaymentReturnMTLS(w http.ResponseWriter, r *http.Request) {
	mtlsClient, ok := h.kaspiClient.(kaspiqr.KaspiQRMTLS)
	if !ok {
		http.Error(w, "Not supported for this client type", http.StatusNotImplemented)
		return
	}

	var req struct {
		DeviceToken string  `json:"deviceToken"`
		QrPaymentID int64   `json:"qrPaymentId"`
		QrReturnID  int64   `json:"qrReturnId"`
		Amount      float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	result, err := mtlsClient.PaymentReturn(ctx, req.DeviceToken, req.QrPaymentID, req.QrReturnID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
