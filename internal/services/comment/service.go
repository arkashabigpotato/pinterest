package comment

import (
	"Project1/internal/models"
	"Project1/internal/repository/comment"
)

type Service interface{
	Create(comment models.Comment) error
	GetByID(commentID int) (*models.Comment, error)
	GetByUserID(userID, limit, offset int) ([]*models.Comment, error)
	GetByPinID(pinID, limit, offset int) ([]*models.Comment, error)
	Delete(commentID int) error
}

type service struct {
	commentRepo      comment.Repository
}

func NewService(commentRepo comment.Repository) Service{
	return &service{
		commentRepo: commentRepo,
	}
}

func (s *service) Create(comment models.Comment) error{
	return s.commentRepo.Create(comment)
}

func (s *service) GetByID(pinID int) (*models.Comment, error){
	return s.commentRepo.GetByID(pinID)
}

func (s *service) GetByUserID(userID, limit, offset int) ([]*models.Comment, error){
	return s.commentRepo.GetByUserID(userID, limit, offset)
}

func (s *service) GetByPinID(pinID, limit, offset int) ([]*models.Comment, error){
	return s.commentRepo.GetByUserID(pinID, limit, offset)
}

func (s *service) Delete(pinID int) error {
	return s.commentRepo.Delete(pinID)
}