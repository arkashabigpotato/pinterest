package users

import (
	"Project1/internal/models"
	"database/sql"
)

type Repository interface {
	Create(user models.User) error
	GetByEmail(email string) (*models.User, error)
}

type UserRepository struct {
	db 			*sql.DB
}

func NewUserRepository(db *sql.DB) Repository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user models.User) error{
	_, err := ur.db.Exec(`insert into users(email, password, is_admin, birth_date, username) 
values ($1, $2, $3, $4, $5)`,
		user.Email, user.Password, user.IsAdmin, user.BirthDate, user.Username,
		)

	return err
}

func (ur *UserRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}

	err := ur.db.QueryRow(`select id, email, password, is_admin, birth_date, username 
from users where email = $1`, email).
		Scan(&user.ID, &user.Email, &user.Password, &user.IsAdmin, &user.BirthDate, &user.Username)

	return user, err
}
