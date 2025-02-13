package handlers

import (
	"core-lnurl/lightning"
	"encoding/json"
	"fmt"
	"net/http"
)

// InvoiceRequest generates an invoice for a username
func InvoiceRequest(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	amountMsats := int64(100000) // Default: 100,000 msats (100 sats)

	if username == "" {
		http.Error(w, "Username required", http.StatusBadRequest)
		return
	}

	// Connect to CLN
	client, err := lightning.NewCLNClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to connect to CLN: %v", err), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Request invoice from CLN
	params := map[string]interface{}{
		"amount_msat": amountMsats,
		"description":  fmt.Sprintf("Payment to %s", username),
	}

	resp, err := client.SendCommand("invoice", params)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create invoice: %v", err), http.StatusInternalServerError)
		return
	}

	// Extract payment request
	bolt11, ok := resp["bolt11"].(string)
	if !ok {
		http.Error(w, "Invalid invoice response", http.StatusInternalServerError)
		return
	}

	// Send response
	response := map[string]interface{}{
		"pr": bolt11,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
