package models

import (
	"go-proj/database"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	mig := database.Conn()

	mig.AutoMigrate(&List{}, &Task{}, &User{})

	DB = mig
}
