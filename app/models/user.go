package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
}

type CreateUserInput struct {
	Email                string `json:"email" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
}

type UpdateUserInput struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}
