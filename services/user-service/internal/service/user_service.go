package service

import (
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/models"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/repository"
)


type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.userRepo.GetUser(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.userRepo.DeleteUser(id)
}