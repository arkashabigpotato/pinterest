package message

import (
	"Project1/internal/models"
	"Project1/internal/repository/message"
	"Project1/internal/services/users"
)

type Service interface {
	Create(message models.Message) error
	Get(userID, chatID, limit, offset int) ([]*models.Message, error)
	Delete(id int) error
}

type service struct {
	messageRepo  message.Repository
	usersService users.Service
}

func NewService(messageRepo message.Repository, usersService users.Service) Service {
	return &service{
		messageRepo:  messageRepo,
		usersService: usersService,
	}
}

func (s *service) Create(message models.Message) error {
	return s.messageRepo.Create(message)
}

func (s *service) Get(userID, chatID, limit, offset int) ([]*models.Message, error) {
	m, err := s.messageRepo.Get(userID, chatID, limit, offset)
	if err != nil {
		return nil, err
	}

	for i, c := range m {
		if c.FromID == userID{
			m[i].IsFromMe = true
		}
		toUser, err := s.usersService.GetByID(c.ToID)
		if err != nil {
			return nil, err
		}
		m[i].ToUsername = toUser.Username

		fromUser, err := s.usersService.GetByID(c.FromID)
		if err != nil {
			return nil, err
		}
		m[i].FromUsername = fromUser.Username
	}
	return m, nil
}

func (s *service) Delete(pinID int) error {
	return s.messageRepo.Delete(pinID)
}
