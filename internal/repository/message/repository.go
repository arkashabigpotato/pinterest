package message

import (
	"Project1/internal/models"
	"database/sql"
)

type Repository interface {
	Create(message models.Message) error
	Get(userID, chatID, limit, offset int) ([]*models.Message, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (mr *repository) Create(message models.Message) error {
	_, err := mr.db.Exec(`insert into message(from_id, to_id, text, date_time) values ($1, $2, $3, $4)`,
		message.FromID, message.ToID, message.Text, message.DateTime,
	)
	return err
}

func (mr *repository) Get(userID, chatID, limit, offset int) ([]*models.Message, error) {
	var messages []*models.Message

	rows, err := mr.db.Query(`select id, from_id, to_id, text, date_time from message 
where (from_id = $1 and to_id = $2) or (from_id = $2 and to_id = $1) order by id desc limit $3 offset $4`, userID, chatID, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		message := &models.Message{}
		err := rows.Scan(&message.ID, &message.FromID, &message.ToID, &message.Text, &message.DateTime)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (mr *repository) Delete(id int) error {
	_, err := mr.db.Exec(`delete from message where id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
