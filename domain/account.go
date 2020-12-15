package domain

import (
	"go-proj/cmd/app_error"
	"go-proj/dto"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Balance int    `gorm:"balance"`
	Account string `gorm:"account"`
}

type AccountRepository interface {
	FindAll() ([]Account, *app_error.AppError)
	FindById(id uint) (*Account, *app_error.AppError)
}

func (a *Account) ToDTO() dto.AccountResponse {
	return dto.AccountResponse{
		ID:      a.ID,
		Balance: a.Balance,
		Account: a.Account,
	}
}
