package service

import (
    "github.com/kollekcioner47/finance-app/internal/models"
    "github.com/kollekcioner47/finance-app/internal/repository"
)

type CategoryService struct {
    repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
    return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(userID int, name, categoryType string) (*models.Category, error) {
    cat := &models.Category{
        UserID: userID,
        Name:   name,
        Type:   categoryType,
    }
    if err := s.repo.Create(cat); err != nil {
        return nil, err
    }
    return cat, nil
}

func (s *CategoryService) GetUserCategories(userID int) ([]*models.Category, error) {
    return s.repo.GetByUserID(userID)
}
