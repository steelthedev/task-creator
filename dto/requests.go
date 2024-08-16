package dto

type TaskCreateRequest struct {
	Title string `json:"title" binding:"required"`
}
