package users

import (
	"Project1/internal/models"
	"database/sql"
)

type Repository interface {
	Create(user models.User) (int, error)
	GetByEmail(email string) (*models.User, error)
	GetAll(limit, offset int) ([]*models.User, error)
	GetByID(userID int) (*models.User, error)
	Update(user models.User) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) Repository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user models.User) (int, error) {
	id := 0
	err := ur.db.QueryRow(`insert into users(email, password, is_admin, birth_date, username, profile_img, status) 
values ($1, $2, $3, $4, $5, $6, $7) returning id`,
		user.Email, user.Password, user.IsAdmin, user.BirthDate, user.Username, user.ProfileImg, user.Status,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ur *UserRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}

	err := ur.db.QueryRow(`select id, email, password, is_admin, birth_date, username, profile_img, status 
from users where email = $1`, email).
		Scan(&user.ID, &user.Email, &user.Password, &user.IsAdmin, &user.BirthDate, &user.Username, &user.ProfileImg, &user.Status)

	return user, err
}

func (ur *UserRepository) GetAll(limit, offset int) ([]*models.User, error) {
	var users []*models.User

	rows, err := ur.db.Query(`select id, email, password, is_admin, birth_date, username, profile_img, status 
from users limit $1 offset $2`, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.IsAdmin, &user.BirthDate, &user.Username, &user.ProfileImg, &user.Status)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) GetByID(userID int) (*models.User, error) {
	user := &models.User{}

	err := ur.db.QueryRow(`select id, email, password, is_admin, birth_date, username, profile_img, status 
from users where id = $1`, userID).
		Scan(&user.ID, &user.Email, &user.Password, &user.IsAdmin, &user.BirthDate, &user.Username, &user.ProfileImg, &user.Status)

	return user, err
}

func (ur *UserRepository) Update(user models.User) error {
	err := ur.db.QueryRow(`update users set status = $1, profile_img = $2 where id = $3`, user.Status, user.ProfileImg, user.ID).Err()
	if err != nil {
		return err
	}
	return nil
}
