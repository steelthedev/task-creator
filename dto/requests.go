package dto

import "time"

type TaskCreateRequest struct {
	Title   string    `json:"title" binding:"required"`
	EndTime time.Time `json:"end_at" binding:"required"`
}

type UserCreateRequest struct {
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name" `
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
