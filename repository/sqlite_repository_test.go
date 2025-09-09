package repository

import (
	"testing"
	"time"
)

func TestSQLite_Migrate(t *testing.T) {
	err := testRepo.Migrate()

	if err != nil {
		t.Error("Migrate failed", err)
	}
}

func TestSQLiteRepository_InsertHolding(t *testing.T) {
	h := Holdings{
		Amount:        1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}

	result, err := testRepo.InsertHolding(h)
	if err != nil {
		t.Error("insert failed", err)
	}

	if result.ID <= 0 {
		t.Error("invalid id sent back")
	}
}

func TestSQLiteRepository_AllHoldings(t *testing.T) {
	h, err := testRepo.AllHoldings()

	if err != nil {
		t.Error("all holdings failed", err)
	}

	if len(h) == 0 {
		t.Error("expected some data")
	}
}
