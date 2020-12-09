package domain

import (
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

func (u UserRepositoryDB) FindById(id uint) (User, error) {
	var user User

	u.DB.Where("id = ?", id).First(&user)

	if (User{}) == user {
		return User{}, fmt.Errorf("User not found")
	}

	return user, nil
}

func NewUserRepositoryDB() UserRepositoryDB {
	return UserRepositoryDB{Conn.Postgres.Conn}
}
