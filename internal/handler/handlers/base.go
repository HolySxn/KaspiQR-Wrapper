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

// GET /health/ping
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling Ping request")
	ctx := r.Context()

	if err := h.kaspiClient.Ping(ctx); err != nil {
		h.logger.Error("Ping failed", slog.String("error", err.Error()))
		http.Error(w, "An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Ping successful")
	w.WriteHeader(http.StatusOK)
}

// GET /partner/tradepoints
func (h *Handler) GetTradePoints(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling GetTradePoints request")
	ctx := r.Context()

	points, err := h.kaspiClient.GetTradePoints(ctx)
	if err != nil {
		h.logger.Error("Failed to get trade points", slog.String("error", err.Error()))
		http.Error(w, "Unable to retrieve trade points at the moment. Please try again later.", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Successfully retrieved trade points")
	json.NewEncoder(w).Encode(points)
}

// POST /device/register
func (h *Handler) DeviceRegister(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling DeviceRegister request")
	var req struct {
		DeviceID     string `json:"deviceId"`
		TradePointID int64  `json:"tradePointId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Invalid request body", slog.String("error", err.Error()))
		http.Error(w, "Invalid request format. Please check your input.", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	token, err := h.kaspiClient.DeviceRegister(ctx, req.DeviceID, req.TradePointID)
	if err != nil {
		h.logger.Error("Failed to register device", slog.String("error", err.Error()))
		http.Error(w, "Unable to register the device at the moment. Please try again later.", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Successfully registered device", slog.String("deviceId", req.DeviceID))
	json.NewEncoder(w).Encode(token)
}

// POST /device/delete
func (h *Handler) DeviceDelete(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling DeviceDelete request")
	var req struct {
		DeviceToken string `json:"deviceToken"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Invalid request body", slog.String("error", err.Error()))
		http.Error(w, "Invalid request format. Please check your input.", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err := h.kaspiClient.DeviceDelete(ctx, req.DeviceToken)
	if err != nil {
		h.logger.Error("Failed to delete device", slog.String("error", err.Error()))
		http.Error(w, "Unable to delete the device at the moment. Please try again later.", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Successfully deleted device", slog.String("deviceToken", req.DeviceToken))
	w.WriteHeader(http.StatusNoContent)
}

// POST /qr/create
func (h *Handler) CreateQR(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling CreateQR request")
	var req struct {
		DeviceToken string  `json:"deviceToken"`
		Amount      float64 `json:"amount"`
		ExternalID  string  `json:"externalId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Invalid request body", slog.String("error", err.Error()))
		http.Error(w, "Invalid request format. Please check your input.", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	qr, err := h.kaspiClient.CreateQR(ctx, req.DeviceToken, req.Amount, req.ExternalID)
	if err != nil {
		h.logger.Error("Failed to create QR", slog.String("error", err.Error()))
		http.Error(w, "Unable to create QR at the moment. Please try again later.", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Successfully created QR", slog.String("externalId", req.ExternalID))
	json.NewEncoder(w).Encode(qr)
}

// POST /qr/create-link
func (h *Handler) CreateLink(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling CreateLink request")
	var req struct {
		DeviceToken string  `json:"deviceToken"`
		Amount      float64 `json:"amount"`
		ExternalID  string  `json:"externalId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Invalid request body", slog.String("error", err.Error()))
		http.Error(w, "Invalid request format. Please check your input.", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	link, err := h.kaspiClient.CreateLink(ctx, req.DeviceToken, req.Amount, req.ExternalID)
	if err != nil {
		h.logger.Error("Failed to create link", slog.String("error", err.Error()))
		http.Error(w, "Unable to create the link at the moment. Please try again later.", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Successfully created link", slog.String("externalId", req.ExternalID))
	json.NewEncoder(w).Encode(link)
}

// GET /payment/status/{qrPaymentId}
func (h *Handler) GetPaymentStatus(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling GetPaymentStatus request")
	vars := mux.Vars(r)
	idStr := vars["QrPaymentId"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.logger.Error("Invalid qrPaymentId", slog.String("qrPaymentId", idStr), slog.String("error", err.Error()))
		http.Error(w, "Invalid QR Payment ID. Please check your input.", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	status, err := h.kaspiClient.GetPaymentStatus(ctx, id)
	if err != nil {
		h.logger.Error("Failed to get payment status", slog.String("error", err.Error()))
		http.Error(w, "Unable to retrieve payment status at the moment. Please try again later.", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Successfully retrieved payment status", slog.Int64("qrPaymentId", id))
	json.NewEncoder(w).Encode(status)
}
