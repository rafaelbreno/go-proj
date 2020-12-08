package domain

import (
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

func NewUserRepositoryDB() UserRepositoryDB {
	return UserRepositoryDB{Conn.Postgres.Conn}
}
