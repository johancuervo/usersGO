package service

import (
	"errors"

	"github.com/johancuervo/usersGO/internal/domain"
	"github.com/johancuervo/usersGO/internal/port"
)

type UserService struct {
	repo port.UserRepository
}

// constructor
func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.repo.Save(user)
}
func (s *UserService) GetAllUsers() ([]domain.User, error) {
	return s.repo.GetAllUsers()
}
func (s *UserService) GetUser(id string) (*domain.User, error) {
	return s.repo.FindByID(id)
}
func (s *UserService) UpdateUser(user *domain.User, id string) error {
	if user.Name == "" {
		return errors.New("name cannot be empty")
	}
	return s.repo.UpdateUser(user, id)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}
