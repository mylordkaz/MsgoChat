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
		INSERT INTO users (username, password, email, google_id, github_id, provider)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	err := r.db.QueryRow(query, user.Username, user.Password, user.Email, user.GoogleID, user.GithubID, user.Provider).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (r *AuthRepository) GetUserByUsername(username string) (*models.User, error){
	query := `
		SELECT id, username, password, email, password, google_id, github_id, provider
		FROM users
		WHERE username = $1
	`
	user := &models.User{}
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.GoogleID, &user.GithubID, &user.Provider)
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
		SELECT id, username, password, email, password, google_id, github_id, provider
		FROM users
		WHERE username = $1
	`
	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.GoogleID, &user.GithubID, &user.Provider)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("erro fetching user: %w", err)
	}
	return user, nil
}

func (r *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, username, password, email, password, google_id, github_id, provider
		FROM users
		WHERE email = $1
	`
	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.GoogleID, &user.GithubID, &user.Provider)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}
	return user, nil
}
func (r *AuthRepository) GetUserByProviderId(provider, providerID string) (*models.User, error) {
    query := `
        SELECT id, username, email, password, google_id, github_id, provider
        FROM users
        WHERE (provider = $1 AND google_id = $2) OR (provider = $1 AND github_id = $2)
    `
    user := &models.User{}
    err := r.db.QueryRow(query, provider, providerID).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.GoogleID, &user.GithubID, &user.Provider)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user not found")
        }
        return nil, fmt.Errorf("error fetching user: %w", err)
    }
    return user, nil
}

func (r *AuthRepository) UpdateUserOAuthInfo(user *models.User) error {
	query := `
		UPDATE users
		SET google_id = $1, github_id = $2
		WHERE id = $3
	`
	_, err := r.db.Exec(query, user.GoogleID, user.GithubID, user.ID)
	if err != nil {
		return fmt.Errorf("error updating user OAuth info: %w", err)
	}
	return nil
}