package repository

import (
	"database/sql"

	"github.com/kollekcioner47/finance-app/internal/models"
)

type transactionRepo struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepo{db: db}
}

func (r *transactionRepo) Create(tx *models.Transaction) error {
	query := `INSERT INTO transactions (user_id, category_id, amount, description, date) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`
	return r.db.QueryRow(query, tx.UserID, tx.CategoryID, tx.Amount, tx.Description, tx.Date).
		Scan(&tx.ID, &tx.CreatedAt)
}

func (r *transactionRepo) GetByID(id int) (*models.Transaction, error) {
	var t models.Transaction
	var categoryID sql.NullInt64
	query := `SELECT id, user_id, category_id, amount, description, date, created_at 
              FROM transactions WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&t.ID, &t.UserID, &categoryID, &t.Amount, &t.Description, &t.Date, &t.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	if categoryID.Valid {
		t.CategoryID = int(categoryID.Int64)
	} else {
		t.CategoryID = 0 // или можно оставить 0, что означает "без категории"
	}
	return &t, nil
}

func (r *transactionRepo) GetByUserID(userID int) ([]*models.Transaction, error) {
	rows, err := r.db.Query(
		`SELECT id, user_id, category_id, amount, description, date, created_at 
         FROM transactions WHERE user_id = $1 ORDER BY date DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		var t models.Transaction
		var categoryID sql.NullInt64
		if err := rows.Scan(
			&t.ID, &t.UserID, &categoryID, &t.Amount, &t.Description, &t.Date, &t.CreatedAt,
		); err != nil {
			return nil, err
		}
		if categoryID.Valid {
			t.CategoryID = int(categoryID.Int64)
		} else {
			t.CategoryID = 0
		}
		transactions = append(transactions, &t)
	}
	return transactions, rows.Err()
}

func (r *transactionRepo) Update(tx *models.Transaction) error {
	query := `UPDATE transactions 
              SET category_id = $1, amount = $2, description = $3, date = $4 
              WHERE id = $5`
	_, err := r.db.Exec(query, tx.CategoryID, tx.Amount, tx.Description, tx.Date, tx.ID)
	return err
}

func (r *transactionRepo) Delete(id int) error {
	query := `DELETE FROM transactions WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
