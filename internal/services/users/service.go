package users

import (
	"Project1/internal/models"
	"Project1/internal/repository/users"
)

type Service interface {
	Create(user models.User) (int, error)
	GetByEmail(email string) (*models.User, error)
	GetByID(userID int) (*models.User, error)
	GetAll(limit, offset int) ([]*models.User, error)
	Update(user models.User) error
}

type service struct {
	usersRepo users.Repository
}

func NewService(usersRepo users.Repository) Service {
	return &service{
		usersRepo: usersRepo,
	}
}

func (s *service) Create(user models.User) (int, error) {
	return s.usersRepo.Create(user)
}

func (s *service) GetByEmail(email string) (*models.User, error) {
	return s.usersRepo.GetByEmail(email)
}

func (s *service) GetByID(userID int) (*models.User, error) {
	return s.usersRepo.GetByID(userID)
}

func (s *service) GetAll(limit, offset int) ([]*models.User, error) {
	return s.usersRepo.GetAll(limit, offset)
}

func (s *service) Update(user models.User) error {
	return s.usersRepo.Update(user)
}
