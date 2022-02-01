package message

import (
	"Project1/internal/models"
	"Project1/internal/repository/message"
)

type Service interface{
	Create(message models.Message) error
	Get(userID, limit, offset int) ([]*models.Message, error)
	Delete(id int) error
}

type service struct {
	messageRepo      message.Repository
}

func NewService(messageRepo message.Repository) Service{
	return &service{
		messageRepo: messageRepo,
	}
}

func (s *service) Create(message models.Message) error{
	return s.messageRepo.Create(message)
}

func (s *service) Get(userID, limit, offset int) ([]*models.Message, error){
	return s.messageRepo.Get(userID, limit, offset)
}

func (s *service) Delete(pinID int) error {
	return s.messageRepo.Delete(pinID)
}