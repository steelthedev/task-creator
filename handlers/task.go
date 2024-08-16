package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/task-handler/services"
	"github.com/steelthedev/task-handler/types"
)

type TaskHandler struct {
	TaskService services.TaskService
}

// Get all tasks handler
// Uses the GetAllTasks services method
func (ts *TaskHandler) GetTasks(ctx *gin.Context) {

	tasks, err := ts.TaskService.GetAllTasks()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, types.ErrResponse{Message: "Internal Error"})
	}
}
