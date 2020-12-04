package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email    string `gorm:"email"`
	Password string `gorm:"password"`
	Status   uint

	// Account can be a uuid string
	Account string `gorm:"account"`
}

type UserRepository interface {
	FindAll() ([]User, error)
}
