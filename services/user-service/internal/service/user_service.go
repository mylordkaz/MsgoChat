package service

import "github.com/mylordkaz/MsgoChat/services/user-service/internal/repository"


type UserService struct {
	userRepo *repository.UserRepository
}
