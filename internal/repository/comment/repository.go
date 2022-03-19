package comment

import (
	"Project1/internal/models"
	"database/sql"
)

type Repository interface{
	Create(comment models.Comment) error
	GetByID(commentID int) (*models.Comment, error)
	GetByUserID(userID, limit, offset int) ([]*models.Comment, error)
	GetByPinID(pinID, limit, offset int) ([]*models.Comment, error)
	Delete(commentID int) error
}

type repository struct {
	db		  *sql.DB
}

func NewCommentRepository(db *sql.DB) Repository  {
	return &repository{
		db: db,
	}
}

func (cr *repository) Create(comment models.Comment) error{
	_, err := cr.db.Exec(`insert into comment(is_deleted, pin_id, text, author_id, date_time) 
values ($1, $2, $3, $4, $5)`,
		comment.IsDeleted, comment.PinID, comment.Text, comment.AuthorID, comment.DateTime,
	)

	return err
}

func (cr *repository) GetByID(commentID int) (*models.Comment, error){
	comment := &models.Comment{}

	err := cr.db.QueryRow(`select id, is_deleted, pin_id, text, author_id, date_time 
from comment where id = $1`, commentID).
		Scan(&comment.ID, &comment.IsDeleted, &comment.PinID, &comment.Text, &comment.AuthorID, &comment.DateTime)

	return comment, err
}

func (cr *repository) GetByUserID(userID, limit, offset int) ([]*models.Comment, error){
	var comments []*models.Comment

	rows, err := cr.db.Query(`select id,  is_deleted, pin_id, text, author_id, date_time from comment 
where author_id = $1 limit $2 offset $3`, userID, limit, offset)
	if err != nil{
		return nil, err
	}

	for rows.Next() {
		comment := &models.Comment{}
		err := rows.Scan(&comment.ID, &comment.IsDeleted, &comment.PinID, &comment.Text, &comment.AuthorID, &comment.DateTime)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (cr *repository) GetByPinID(pinID, limit, offset int) ([]*models.Comment, error) {
	var comments []*models.Comment

	rows, err := cr.db.Query(`select id,  is_deleted, pin_id,  text, author_id, date_time 
from comment where pin_id = $1 limit $2 offset $3`, pinID, limit, offset)
	if err != nil{
		return nil, err
	}

	for rows.Next() {
		comment := &models.Comment{}
		err := rows.Scan(&comment.ID, &comment.IsDeleted, &comment.PinID, &comment.Text, &comment.AuthorID, &comment.DateTime)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (cr *repository) Delete(commentID int) error {
	_, err := cr.db.Exec(`delete from comment where id = $1`, commentID)
	if err != nil {
		return err
	}
	return nil
}
