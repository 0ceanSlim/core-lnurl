package server

import (
	"fmt"
	"log"
	"net/http"

	"core-lnurl/config"
	"core-lnurl/handlers"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()

	// LNURL-p endpoint for well-known URL
	router.HandleFunc("/.well-known/lnurlp/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handlers.LNURLpHandler(w, r)
	})
	
	// Original LNURL-pay endpoint with query parameters
	router.HandleFunc("/lnurl/pay", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handlers.InvoiceRequest(w, r)
	})
	
	// NEW: Path-based LNURL-pay endpoint for better client compatibility
	// This will handle patterns like /lnurl/pay/username
	router.HandleFunc("/lnurl/pay/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handlers.PathInvoiceRequest(w, r)
	})

	port := config.AppConfig.Server.Port
	fmt.Printf("LNURL Server running on port %d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}