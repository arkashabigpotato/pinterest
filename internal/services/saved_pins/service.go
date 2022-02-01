package saved_pins

import (
	"Project1/internal/models"
	"Project1/internal/repository/saved_pins"
)

type Service interface{
	Append(savedPin models.SavedPin) error
	GetByUserID(userID, limit, offset int) ([]*models.SavedPin, error)
	Delete(pinID int) error
}

type service struct {
	savedPinRepo      saved_pins.Repository
}

func NewService(savedPinRepo saved_pins.Repository) Service{
	return &service{
		savedPinRepo: savedPinRepo,
	}
}

func (s *service) Append(savedPin models.SavedPin) error{
	return s.savedPinRepo.Append(savedPin)
}

func (s *service) GetByUserID(userID, limit, offset int) ([]*models.SavedPin, error){
	return s.savedPinRepo.GetByUserID(userID, limit, offset)
}

func (s *service) Delete(pinID int) error {
	return s.savedPinRepo.Delete(pinID)
}