package repository

import "github.com/kollekcioner47/finance-app/internal/models"

type CategoryRepository interface {
    Create(category *models.Category) error
    GetByID(id int) (*models.Category, error)
    GetByUserID(userID int) ([]*models.Category, error)
    Update(category *models.Category) error
    Delete(id int) error
}
