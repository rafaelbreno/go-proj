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

// For store method
type CreateListInput struct {
	UserId uint   `json:"user_id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Status uint   `json:"status" binding:"required"`
}

// For update method
type UpdateListInput struct {
	Title  string `json:"title"`
	Status uint   `json:"status"`
}
