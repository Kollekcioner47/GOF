package repository

import (
	"database/sql"

	"github.com/kollekcioner47/finance-app/internal/models"
)

type categoryRepo struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) Create(category *models.Category) error {
	query := `INSERT INTO categories (user_id, name, type) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRow(query, category.UserID, category.Name, category.Type).Scan(&category.ID)
}

func (r *categoryRepo) GetByID(id int) (*models.Category, error) {
	var cat models.Category
	query := `SELECT id, user_id, name, type FROM categories WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&cat.ID, &cat.UserID, &cat.Name, &cat.Type)
	if err != nil {
		return nil, err
	}
	return &cat, nil
}

func (r *categoryRepo) GetByUserID(userID int) ([]*models.Category, error) {
	rows, err := r.db.Query(`SELECT id, user_id, name, type FROM categories WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.UserID, &cat.Name, &cat.Type); err != nil {
			return nil, err
		}
		categories = append(categories, &cat)
	}
	return categories, rows.Err()
}

func (r *categoryRepo) Update(category *models.Category) error {
	query := `UPDATE categories SET name = $1, type = $2 WHERE id = $3`
	_, err := r.db.Exec(query, category.Name, category.Type, category.ID)
	return err
}

func (r *categoryRepo) Delete(id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
