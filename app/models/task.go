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
type Task struct {
	gorm.Model
	ListId uint
	Title  string
	Status uint
}

// For store method
type CreateTaskInput struct {
	ListId uint   `json:"list_id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Status uint   `json:"status" binding:"required"`
}

// For update method
type UpdateTaskInput struct {
	Title  string `json:"title"`
	Status uint   `json:"status"`
}
