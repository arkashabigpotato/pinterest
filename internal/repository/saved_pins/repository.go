package saved_pins

import (
	"Project1/internal/models"
	"database/sql"
)

type Repository interface {
	Append(savedPin models.SavedPin) error
	GetByUserID(userID, limit, offset int) ([]*models.SavedPin, error)
	Delete(pinID int) error
}

type repository struct {
	db *sql.DB
}

func NewSavedPinRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (sr *repository) Append(savedPin models.SavedPin) error {
	_, err := sr.db.Exec(`insert into saved_pins(pin_id, user_id) values ($1, $2)`,
		savedPin.PinID, savedPin.UserID,
	)

	return err
}

func (sr *repository) GetByUserID(userID, limit, offset int) ([]*models.SavedPin, error) {
	var savedPins []*models.SavedPin

	rows, err := sr.db.Query(`select user_id, pin_id from saved_pins 
where user_id = $1  limit $2 offset $3`, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		savedPin := &models.SavedPin{}
		err := rows.Scan(&savedPin.UserID, &savedPin.PinID)
		if err != nil {
			return nil, err
		}
		savedPins = append(savedPins, savedPin)
	}
	return savedPins, nil
}

func (sr *repository) Delete(pinID int) error {
	_, err := sr.db.Exec(`delete from saved_pins where pin_id = $1`, pinID)
	if err != nil {
		return err
	}
	return nil
}
