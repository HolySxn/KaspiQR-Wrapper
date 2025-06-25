package httpHandler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	kaspiqr "github.com/HolySxn/KaspiQR-Wrapper/internal/kaspi-qr"
)

type Handler struct {
	logger      *slog.Logger
	kaspiClient kaspiqr.KaspiQRBase
}

func NewHandler(logger *slog.Logger, kaspiClient kaspiqr.KaspiQRBase) *Handler {
	return &Handler{
		logger:      logger,
		kaspiClient: kaspiClient,
	}
}

// GET /partner/tradepoints
func (h *Handler) GetTradePoints(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	points, err := h.kaspiClient.GetTradePoints(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(points)
}

// POST /device/register
func (h *Handler) DeviceRegister(w http.ResponseWriter, r *http.Request) {
	var req struct {
		DeviceID     string `json:"deviceId"`
		TradePointID int64  `json:"tradePointId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	token, err := h.kaspiClient.DeviceRegister(ctx, req.DeviceID, req.TradePointID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(token)
}

// POST /device/delete
func (h *Handler) DeviceDelete(w http.ResponseWriter, r *http.Request) {
	var req struct {
		DeviceToken string `json:"deviceToken"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err := h.kaspiClient.DeviceDelete(ctx, req.DeviceToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// POST /qr/create
func (h *Handler) CreateQR(w http.ResponseWriter, r *http.Request) {
	var req struct {
		DeviceToken string  `json:"deviceToken"`
		Amount      float64 `json:"amount"`
		ExternalID  string  `json:"externalId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	qr, err := h.kaspiClient.CreateQR(ctx, req.DeviceToken, req.Amount, req.ExternalID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(qr)
}

// POST /qr/create-link
func (h *Handler) CreateLink(w http.ResponseWriter, r *http.Request) {
	var req struct {
		DeviceToken string  `json:"deviceToken"`
		Amount      float64 `json:"amount"`
		ExternalID  string  `json:"externalId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	link, err := h.kaspiClient.CreateLink(ctx, req.DeviceToken, req.Amount, req.ExternalID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(link)
}

// GET /payment/status/{qrPaymentId}
func (h *Handler) GetPaymentStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["QrPaymentId"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid qrPaymentId", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	status, err := h.kaspiClient.GetPaymentStatus(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(status)
}
