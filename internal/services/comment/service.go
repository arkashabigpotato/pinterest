package comment

import (
	"Project1/internal/models"
	"Project1/internal/repository/comment"
	"Project1/internal/services/users"
)

type Service interface {
	Create(comment models.Comment) error
	GetByID(commentID int) (*models.Comment, error)
	GetByUserID(userID, limit, offset int) ([]*models.Comment, error)
	GetByPinID(pinID, limit, offset int) ([]*models.Comment, error)
	Delete(commentID int) error
}

type service struct {
	commentRepo  comment.Repository
	usersService users.Service
}

func NewService(commentRepo comment.Repository, usersService users.Service) Service {
	return &service{
		commentRepo:  commentRepo,
		usersService: usersService,
	}
}

func (s *service) Create(comment models.Comment) error {
	return s.commentRepo.Create(comment)
}

func (s *service) GetByID(commentID int) (*models.Comment, error) {
	return s.commentRepo.GetByID(commentID)
}

func (s *service) GetByUserID(userID, limit, offset int) ([]*models.Comment, error) {
	return s.commentRepo.GetByUserID(userID, limit, offset)
}

func (s *service) GetByPinID(pinID, limit, offset int) ([]*models.Comment, error) {
	com, err := s.commentRepo.GetByPinID(pinID, limit, offset)
	if err != nil {
		return nil, err
	}

	for i, c := range com {
		user, err := s.usersService.GetByID(c.AuthorID)
		if err != nil {
			return nil, err
		}
		com[i].Username = user.Username
	}
	return com, nil
}

func (s *service) Delete(commentID int) error {
	return s.commentRepo.Delete(commentID)
}
