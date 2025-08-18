package handlers

import (
	"encoding/json"
	"net/http"
	"pro_mux/internal/constants"
	"pro_mux/internal/entities"
	"pro_mux/internal/repositories"
	"strconv"

	"github.com/gorilla/mux"
)

type WalletHandler struct {
	Repo *repositories.WalletRepository
}

// ----- Auth -----
func checkAuth(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		return false
	}
	return username == constants.AdminName && password == constants.AdminPass
}

func WithAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !checkAuth(r) {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		handler(w, r)
	}
}

// ----- CRUD -----
func (h *WalletHandler) GetWallets(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.Repo.GetAll())
}

func (h *WalletHandler) GetWalletByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	wallet, err := h.Repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, wallet)
}

func (h *WalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var wallet entities.Wallet
	if err := json.NewDecoder(r.Body).Decode(&wallet); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	created := h.Repo.Create(wallet)
	writeJSON(w, http.StatusCreated, created)
}

func (h *WalletHandler) UpdateWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var wallet entities.Wallet
	if err := json.NewDecoder(r.Body).Decode(&wallet); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	updated, err := h.Repo.Update(id, wallet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

func (h *WalletHandler) DeleteWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

// ----- Utils -----
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
