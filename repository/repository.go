package repository

import (
	"errors"
	"time"
)

var (
	errUpdateFailed = errors.New("Update failed")
	errDeleteFailed = errors.New("Delete failed")
)

type Holdings struct {
	ID            int64     `json:"id"`
	Amount        int       `json:"amount"`
	PurchaseDate  time.Time `json:"purchase_date"`
	PurchasePrice int       `json:"purchase_price"`
}

type Repository interface {
	Migrate() error
	InsertHolding(h Holdings) (*Holdings, error)
	AllHoldings([]Holdings, error)
	GetHoldingById(id int) (*Holdings, error)
	UpdateHolding(id int64, updated Holdings) error
	DeleteHolding(id int64) error
}
