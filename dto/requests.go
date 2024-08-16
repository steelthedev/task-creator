package dto

import "time"

type TaskCreateRequest struct {
	Title   string    `json:"title" binding:"required"`
	EndTime time.Time `json:"end_at" binding:"required"`
}
