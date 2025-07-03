package models

import "time"

type Task struct {
	ID        *uint64    `gorm:"primaryKey" json:"id"`
	Title     *string    `json:"title"`
	Content   *string    `json:"content"`
	Deadline  *time.Time `json:"deadline"`
	Done      *bool      `json:"done"`
	CreatedAt *time.Time `json:"created_at"`
}

func (Task) TableName() string { return "tasks" }

type TaskRequest struct {
	Title    string     `json:"title" binding:"required"`
	Content  string     `json:"content" binding:"required"`
	Deadline *time.Time `json:"deadline"`
}

type TaskUpdateRequest struct {
	Title    *string    `json:"title"`
	Content  *string    `json:"content"`
	Deadline *time.Time `json:"deadline"`
	Done     *bool      `json:"done"`
}
