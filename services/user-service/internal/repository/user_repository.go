package repository

import (
	"database/sql"

	"github.com/mylordkaz/MsgoChat/services/user-service/internal/models"
)


type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	query := `
		INSERT INTO users (id, username, email, avatar_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, username, email, avatar_url, created_at, updated_at
	`

	err := r.db.QueryRow(query,
		user.ID, user.Username, user.Email, user.AvatarURl, user.CreatedAt, user.UpdatedAt,
	).Scan(
		&user.ID, &user.Username, &user.Email, &user.AvatarURl, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUser(id string) (*models.User, error) {
	query := `
		SELECT id, username, email, avatarURL, created_at, updated_at
		FROM users WHERE id = $1
	`

	var user models.User
	err :=r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.AvatarURl, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.User)  error {
	query := `
		UPDATE users
		SET username = $2, email = $3, avatar_url $4, updated_at = $5
		WHERE id = $1
	`
	_, err := r.db.Exec(query,
		user.ID, user.Username, user.Email, user.AvatarURl, user.UpdatedAt,
	)
	return err
}

func (r *UserRepository) DeleteUser(id string) error {
	query := `
		DELETE FROM users WHERE id = $1
	`
	_, err := r.db.Exec(query, id)

	return err
}
