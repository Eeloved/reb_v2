package main

import (
	"fmt"
	"log"
	"net/http"
	"pro_mux/internal/handlers"
	"pro_mux/internal/repositories"

	"github.com/gorilla/mux"
)

func main() {
	repo := repositories.NewWalletRepository()
	handler := &handlers.WalletHandler{Repo: repo}

	r := mux.NewRouter()

	r.HandleFunc("/wallets", handlers.WithAuth(handler.GetWallets)).Methods("GET")
	r.HandleFunc("/wallets/{id:[0-9]+}", handlers.WithAuth(handler.GetWalletByID)).Methods("GET")
	r.HandleFunc("/wallets", handlers.WithAuth(handler.CreateWallet)).Methods("POST")
	r.HandleFunc("/wallets/{id:[0-9]+}", handlers.WithAuth(handler.UpdateWallet)).Methods("PUT")
	r.HandleFunc("/wallets/{id:[0-9]+}", handlers.WithAuth(handler.DeleteWallet)).Methods("DELETE")

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
