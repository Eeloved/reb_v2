package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// ----- сущность -----
type Wallet struct {
	Name    string  `json:"name"`
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
}

// ----- "база данных" (в памяти) -----
var (
	wallets = make(map[int]Wallet)
	nextID  = 1
	mu      sync.Mutex
)

// ----- хендлеры -----

// список всех кошельков
func getWallets(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var list []Wallet
	for _, v := range wallets {
		list = append(list, v)
	}
	writeJSON(w, http.StatusOK, list)
}

// получение одного кошелька по id (?id=1)
func getWalletByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	wallet, ok := wallets[id]
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, wallet)
}

// создание нового кошелька
func createWallet(w http.ResponseWriter, r *http.Request) {
	var newWallet Wallet
	if err := json.NewDecoder(r.Body).Decode(&newWallet); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	newWallet.ID = nextID
	nextID++
	wallets[newWallet.ID] = newWallet

	writeJSON(w, http.StatusCreated, newWallet)
}

// удаление кошелька (?id=1)
func deleteWallet(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, ok := wallets[id]; !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	delete(wallets, id)
	writeJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

// ----- утилита для JSON -----
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// ----- запуск сервера -----
func main() {
	http.HandleFunc("/wallets", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// если есть query id → вернуть один
			if r.URL.Query().Get("id") != "" {
				getWalletByID(w, r)
			} else {
				getWallets(w, r)
			}
		case http.MethodPost:
			createWallet(w, r)
		case http.MethodDelete:
			deleteWallet(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// $ curl -X POST -d '{"name": "Jhon", "balance":900}' -H "Content-Type: application/json" http://localhost:8080/wallets

// $ curl http://localhost:8080/wallets?id=1

// $ curl -X DELETE http://localhost:8080/wallets?id=1

// $ curl http://localhost:8080/wallets
