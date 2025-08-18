package repositories

import (
	"errors"
	"pro_mux/internal/entities"
	"sync"
)

type WalletRepository struct {
	data   map[int]entities.Wallet
	nextID int
	mu     sync.Mutex
}

func NewWalletRepository() *WalletRepository {
	return &WalletRepository{
		data:   make(map[int]entities.Wallet),
		nextID: 1,
	}
}

func (r *WalletRepository) GetAll() []entities.Wallet {
	r.mu.Lock()
	defer r.mu.Unlock()

	var wallets []entities.Wallet
	for _, w := range r.data {
		wallets = append(wallets, w)
	}
	return wallets
}

func (r *WalletRepository) GetByID(id int) (entities.Wallet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	w, ok := r.data[id]
	if !ok {
		return entities.Wallet{}, errors.New("wallet not found")
	}
	return w, nil
}

func (r *WalletRepository) Create(wallet entities.Wallet) entities.Wallet {
	r.mu.Lock()
	defer r.mu.Unlock()

	wallet.ID = r.nextID
	r.nextID++
	r.data[wallet.ID] = wallet
	return wallet
}

func (r *WalletRepository) Update(id int, wallet entities.Wallet) (entities.Wallet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return entities.Wallet{}, errors.New("wallet not found")
	}

	wallet.ID = id
	r.data[id] = wallet
	return wallet, nil
}

func (r *WalletRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return errors.New("wallet not found")
	}
	delete(r.data, id)
	return nil
}
