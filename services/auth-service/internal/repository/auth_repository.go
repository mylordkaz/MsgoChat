package repository

import (
	"database/sql"
	"fmt"

	"github.com/mylordkaz/MsgoChat/services/auth-service/internal/models"
)

type AuthRepository struct {
	db *sql.DB
}
func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (username, password, email)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := r.db.QueryRow(query, user.Username, user.Password, user.Email).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (r *AuthRepository) GetUserByUsername(username string) (*models.User, error){
	query := `
		SELECT id, username, password, email
		FROM users
		WHERE username = $1
	`
	user := &models.User{}
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}
	return user, nil
}

func (r *AuthRepository) GetUserByID(id int) (*models.User, error){
	query := `
		SELECT id, username, password, email
		FROM users
		WHERE username = $1
	`
	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("erro fetching user: %w", err)
	}
	return user, nil
}