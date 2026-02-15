package repository

import (
	"database/sql"

	"github.com/kollekcioner47/finance-app/internal/models"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *models.User) error {
	query := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id, created_at`
	return r.db.QueryRow(query, user.Email, user.PasswordHash).Scan(&user.ID, &user.CreatedAt)
}

func (r *userRepo) GetByID(id int) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password_hash, created_at FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password_hash, created_at FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Update(user *models.User) error {
	query := `UPDATE users SET email = $1, password_hash = $2 WHERE id = $3`
	_, err := r.db.Exec(query, user.Email, user.PasswordHash, user.ID)
	return err
}

func (r *userRepo) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
