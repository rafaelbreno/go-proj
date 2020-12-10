package domain

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepositoryDB struct {
	DB *gorm.DB
}

func (u UserRepositoryDB) FindAll() ([]User, error) {
	var users []User

	u.DB.Find(&users)

	return users, nil
}

func (u UserRepositoryDB) FindById(id uint) (*User, error) {
	var user User

	if err := u.DB.Where("id = ?", id).First(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &User{}, fmt.Errorf("Unable to stablish a DB connection")
	}

	if (User{}) == user {
		return &User{}, fmt.Errorf("User not found")
	}

	return &user, nil
}

func NewUserRepositoryDB() UserRepositoryDB {
	return UserRepositoryDB{Conn.Postgres.Conn}
}
