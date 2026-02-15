package repository

import "github.com/kollekcioner47/finance-app/internal/models"

type TransactionRepository interface {
    Create(tx *models.Transaction) error
    GetByID(id int) (*models.Transaction, error)
    GetByUserID(userID int) ([]*models.Transaction, error)
    Update(tx *models.Transaction) error
    Delete(id int) error
}
