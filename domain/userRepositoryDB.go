package domain

import (
	"errors"

	"go-proj/cmd/app_error"

	"gorm.io/gorm"
)

type UserRepositoryDB struct {
	DB *gorm.DB
}

func (u UserRepositoryDB) FindAll() ([]User, *app_error.AppError) {
	var users []User

	u.DB.Find(&users)

	return users, nil
}

func (u UserRepositoryDB) FindById(id uint) (*User, *app_error.AppError) {
	var user User

	if err := u.DB.Where("id = ?", id).First(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &User{}, app_error.NewUnexpectedError("Unable to stablish a DB connection", "domain/userRepositoryDB/FindById")
	}

	if (User{}) == user {
		return &User{}, app_error.NewNotFoundError("UserNotFound", "domain/userRepositoryDB/FindById")
	}

	return &user, nil
}

func NewUserRepositoryDB() UserRepositoryDB {
	return UserRepositoryDB{Conn.Postgres.Conn}
}
