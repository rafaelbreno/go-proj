package seeder

import (
	"go-proj/domain"

	"github.com/bxcodec/faker/v3"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Seed struct {
	db *gorm.DB
}

// Setting a Seed instance
func Init(db *gorm.DB) {
	var s Seed
	s.db = db
	if db.Find(&domain.User{}).RowsAffected == 0 {
		s.UserSeeder(10)
	}
}

func (s *Seed) UserSeeder(quant int) *Seed {
	for i := 0; i < quant; i++ {
		status, _ := faker.RandomInt(0, 2)
		account, _ := uuid.NewV4()
		user := domain.User{
			Email:    faker.Email(),
			Password: faker.Password(),
			Status:   uint(status[0]),
			Account:  account.String(),
		}

		if res := s.db.Create(&user); res.Error != nil {
			panic(res.Error)
		}
	}

	return s
}
