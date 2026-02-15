package service

import (
    "errors"
    "github.com/kollekcioner47/finance-app/internal/models"
    "github.com/kollekcioner47/finance-app/internal/repository"
    "time"
)

type TransactionService struct {
    repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
    return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(userID, categoryID int, amount float64, description string, date time.Time) (*models.Transaction, error) {
    if amount <= 0 {
        return nil, errors.New("amount must be positive")
    }
    tx := &models.Transaction{
        UserID:      userID,
        CategoryID:  categoryID,
        Amount:      amount,
        Description: description,
        Date:        date,
    }
    if err := s.repo.Create(tx); err != nil {
        return nil, err
    }
    return tx, nil
}

func (s *TransactionService) GetUserTransactions(userID int) ([]*models.Transaction, error) {
    return s.repo.GetByUserID(userID)
}
