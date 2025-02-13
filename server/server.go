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

	// LNURLp (well-known LNURL payment endpoints)
	router.HandleFunc("/.well-known/lnurlp/{username}", handlers.LNURLpHandler).Methods("GET")

	// LNURL-Pay (handles invoice generation)
	router.HandleFunc("/lnurl/pay", handlers.InvoiceRequest).Methods("GET")

	port := config.AppConfig.Server.Port
	fmt.Printf("LNURL Server running on port %d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}