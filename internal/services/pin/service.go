package pin

import (
	"Project1/internal/models"
	"Project1/internal/repository/pin"

)

type Service interface {
	Create(pin models.Pin) error
	GetByUserID(userID, limit, offset int) ([]*models.Pin, error)
	GetByID(pinID int) (*models.Pin, error)
	GetAll(limit, offset int) ([]*models.Pin, error)
	Delete(pinID int) error
}

type service struct {
	pinRepo      pin.Repository
}

func NewService(pinRepo pin.Repository) Service{
	return &service{
		pinRepo: pinRepo,
	}
}

func (s *service) Create(pin models.Pin) error{
	return s.pinRepo.Create(pin)
}

func (s *service) GetByID(pinID int) (*models.Pin, error){
	return s.pinRepo.GetByID(pinID)
}

func (s *service) GetByUserID(userID, limit, offset int) ([]*models.Pin, error){
	return s.pinRepo.GetByUserID(userID, limit, offset)
}

func (s *service) GetAll(limit, offset int) ([]*models.Pin, error){
	return s.pinRepo.GetAll(limit, offset)
}

func (s *service) Delete(pinID int) error {
	return s.pinRepo.Delete(pinID)
}