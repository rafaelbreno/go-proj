package domain

import (
	"gorm.io/gorm"
)

type UserRepositoryStub struct {
	users []User
}

func (u UserRepositoryStub) FindAll() ([]User, error) {
	return u.users, nil
}

func NewUserRepositoryStub() UserRepositoryStub {
	users := []User{
		{gorm.Model{}, "a", "b", 1, "c"},
	}

	return UserRepositoryStub{users}
}
