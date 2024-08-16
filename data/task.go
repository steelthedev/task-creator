package data

import (
	"time"
)

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"title" gorm:"column:title"`
	CreatedAt time.Time `json:"created_at"`
	EndAt     time.Time `json:"end_at"`
}
