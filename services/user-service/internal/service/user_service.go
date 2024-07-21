package service

import (
	"time"

	"github.com/mylordkaz/MsgoChat/services/user-service/internal/models"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/repository"
	"github.com/mylordkaz/MsgoChat/services/user-service/pkg/utils"
)


type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	// generate new ID
	id, err := utils.GenerateRandomID()
	if err != nil {
		return nil, err
	}
	user.ID = id

	// timestamp
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// validate data
	if err := utils.ValidateUser(user); err != nil {
		return nil, err
	}
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.userRepo.GetUser(id)
}

func (s *UserService) UpdateUser(user *models.User) error {

	if err := utils.ValidateUser(user); err != nil {
		return err
	}

	user.UpdatedAt = time.Now()
	return s.userRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.userRepo.DeleteUser(id)
}

func (s *UserService) ListUsers(page, pageSize int) ([]*models.User, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	return s.userRepo.ListUsers(pageSize, offset)
}