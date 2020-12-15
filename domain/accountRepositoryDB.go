package domain

import (
	"errors"

	"go-proj/cmd/app_error"

	"gorm.io/gorm"
)

type AccountRepositoryDB struct {
	DB *gorm.DB
}

func (u AccountRepositoryDB) FindAll() ([]Account, *app_error.AppError) {
	var accounts []Account

	if err := u.DB.Find(&accounts).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return accounts, app_error.NewUnexpectedError("Unable to stablish a DB connection", "domain/accountRepositoryDB/FindById")
	}

	return accounts, nil
}

func (u AccountRepositoryDB) FindById(id uint) (*Account, *app_error.AppError) {
	var account Account

	if err := u.DB.Where("id = ?", id).First(&account).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &Account{}, app_error.NewUnexpectedError("Unable to stablish a DB connection", "domain/accountRepositoryDB/FindById")
	}

	if (Account{}) == account {
		return &Account{}, app_error.NewNotFoundError("AccountNotFound", "domain/accountRepositoryDB/FindById")
	}

	return &account, nil
}

func NewAccountRepositoryDB() AccountRepositoryDB {
	return AccountRepositoryDB{Conn.Postgres.Conn}
}
