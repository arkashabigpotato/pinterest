package pin

import (
	"Project1/internal/models"
	"database/sql"
)

type Repository interface {
	Create(pin models.Pin) error
	GetByUserID(userID, limit, offset int) ([]*models.Pin, error)
	GetByID(pinID int) (*models.Pin, error)
	GetAll(limit, offset int) ([]*models.Pin, error)
	Delete(pinID int) error
}

type repository struct {
	db		  *sql.DB
}

func NewPinRepository(db *sql.DB) Repository  {
	return &repository{
		db: db,
	}
}

func (pr *repository) Create(pin models.Pin) error{
	_, err := pr.db.Exec(`insert into pin(description, likes_count, dislikes_count, author_id, pin_link) values ($1, $2, $3, $4, $5)`,
		pin.Description, pin.LikesCount, pin.DislikesCount, pin.AuthorID, pin.PinLink,
	)

	return err
}

func (pr *repository) GetByUserID(userID, limit, offset int) ([]*models.Pin, error){
	var pins []*models.Pin

	rows, err := pr.db.Query(`select id, description, likes_count, dislikes_count, author_id, pin_link 
from pin where author_id = $1 limit $2 offset $3`, userID, limit, offset)
	if err != nil{
		return nil, err
	}

	for rows.Next() {
		pin := &models.Pin{}
		err := rows.Scan(&pin.ID, &pin.Description, &pin.LikesCount, &pin.DislikesCount, &pin.AuthorID, &pin.PinLink)
		if err != nil {
			return nil, err
		}
		pins = append(pins, pin)
	}
	return pins, nil
}

func (pr *repository) GetByID(pinID int) (*models.Pin, error){
	pin := &models.Pin{}

	err := pr.db.QueryRow(`select id, description, likes_count, dislikes_count, author_id, pin_link 
from pin where id = $1`, pinID).
		Scan(&pin.ID, &pin.Description, &pin.LikesCount, &pin.DislikesCount, &pin.AuthorID, &pin.PinLink)

	return pin, err
}

func (pr *repository) GetAll(limit, offset int) ([]*models.Pin, error){
	var pins []*models.Pin

	rows, err := pr.db.Query(`select id, description, likes_count, dislikes_count, author_id, pin_link 
from pin limit $1 offset $2`, limit, offset)
	if err != nil{
		return nil, err
	}

	for rows.Next() {
		pin := &models.Pin{}
		err := rows.Scan(&pin.ID, &pin.Description, &pin.LikesCount, &pin.DislikesCount, &pin.AuthorID, &pin.PinLink)
		if err != nil {
			return nil, err
		}
		pins = append(pins, pin)
	}
	return pins, nil
}

func (pr *repository) Delete(pinID int) error{
	_, err := pr.db.Exec(`delete from pin where id = $1`, pinID)
	if err != nil{
		return err
	}
	return nil
}