package models

import (
	"gorm.io/gorm"
)

/* gorm.Model automatically adds:
 * ID        uint
 * CreatedAt time.Time
 * UpdatedAt time.Time
 * DeletedAt time.Time
 **/
type List struct {
	gorm.Model
	UserId uint
	Title  string
	Status uint
}

type ListInput struct {
	UserId uint   `json:"user_id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Status uint   `json:"status" binding:"required"`
}

type ListRepository interface {
	All() (*[]List, error)
	FindById(id uint) (*List, error)
	Store(list *List) (*List, error)
	Update(list *List, id uint) (*List, error)
	Delete(list *List) error
}
