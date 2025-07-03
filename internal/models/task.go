package models

import "time"

type Task struct {
	ID        *uint      `gorm:"primaryKey" json:"id"`
	Title     *string    `json:"title"`
	Content   *string    `json:"content"`
	Deadline  *time.Time `json:"deadline"`
	Done      *bool      `json:"done"`
	CreatedAt *time.Time `json:"created_at"`
}
