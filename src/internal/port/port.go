package port

import "github.com/johancuervo/usersGO/internal/domain"

type UserRepository interface {
	GetAllUsers() ([]domain.User, error)
	FindByID(id string) (*domain.User, error)
	Save(user *domain.User) error
	UpdateUser(user *domain.User, id string) error
	Delete(id string) error
}
