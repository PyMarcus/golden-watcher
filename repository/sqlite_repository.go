package repository

import (
	"database/sql"
	"errors"
	"time"
)

type SQLiteRepository struct {
	Conn *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		Conn: db,
	}
}

func (repo *SQLiteRepository) Migrate() error {
	query := `
		CREATE TABLE IF NOT EXISTS holdings(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			amount REAL NOT NULL,
			purchase_date INTEGER NOT NULL,
			purchase_price INTEGER NOT NULL
		)
	`
	_, err := repo.Conn.Exec(query)
	return err
}

// InsertHolding insere um novo registro na tabela holdings
func (repo *SQLiteRepository) InsertHolding(h Holdings) (*Holdings, error) {
	query := `
		INSERT INTO holdings(amount, purchase_date, purchase_price)
		VALUES(?, ?, ?)
	`
	res, err := repo.Conn.Exec(query, h.Amount, h.PurchaseDate.Unix(), h.PurchasePrice)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	h.ID = id

	return &h, nil
}

// AllHoldings retorna todos os registros
func (repo *SQLiteRepository) AllHoldings() ([]Holdings, error) {
	query := `SELECT id, amount, purchase_date, purchase_price FROM holdings`

	rows, err := repo.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var holdings []Holdings
	for rows.Next() {
		var h Holdings
		var ts int64
		err := rows.Scan(&h.ID, &h.Amount, &ts, &h.PurchasePrice)
		if err != nil {
			return nil, err
		}
		h.PurchaseDate = time.Unix(ts, 0)
		holdings = append(holdings, h)
	}

	return holdings, nil
}

// GetHoldingById retorna um registro pelo ID
func (repo *SQLiteRepository) GetHoldingById(id int) (*Holdings, error) {
	query := `SELECT id, amount, purchase_date, purchase_price FROM holdings WHERE id = ? LIMIT 1`
	row := repo.Conn.QueryRow(query, id)

	var h Holdings
	var ts int64
	err := row.Scan(&h.ID, &h.Amount, &ts, &h.PurchasePrice)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	h.PurchaseDate = time.Unix(ts, 0)

	return &h, nil
}

// UpdateHolding atualiza um registro existente
func (repo *SQLiteRepository) UpdateHolding(id int64, updated Holdings) error {
	query := `
		UPDATE holdings 
		SET amount = ?, purchase_date = ?, purchase_price = ?
		WHERE id = ?
	`
	res, err := repo.Conn.Exec(query, updated.Amount, updated.PurchaseDate.Unix(), updated.PurchasePrice, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errUpdateFailed
	}

	return nil
}

// DeleteHolding remove um registro pelo ID
func (repo *SQLiteRepository) DeleteHolding(id int64) error {
	query := `DELETE FROM holdings WHERE id = ?`
	res, err := repo.Conn.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errDeleteFailed
	}

	return nil
}
