package service

import (
	"github.com/teakingwang/cursor-demo/internal/models"
	"github.com/teakingwang/cursor-demo/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(),
	}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepo.Create(user)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
