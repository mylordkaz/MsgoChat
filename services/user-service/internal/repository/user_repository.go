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
	// implement db insertion logic
}
func (r *UserRepository) GetUser(id string) (*models.User, error) {
	// implement db retrival logic
}
func (r *UserRepository) UpdateUser(user *models.User)  error {
	// implement db update logic
}
func (r *UserRepository) DeleteUser(id string) error {
	// implement db deletion logic
}