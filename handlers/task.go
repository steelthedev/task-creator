package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/task-handler/data"
	"github.com/steelthedev/task-handler/dto"
	"github.com/steelthedev/task-handler/errors"
	"github.com/steelthedev/task-handler/services"
)

type TaskHandler struct {
	TaskService services.TaskService
}

func NewTaskHandler(taskService services.TaskService) *TaskHandler {
	return &TaskHandler{
		TaskService: taskService,
	}
}

// Get all tasks handler
// Uses the GetAllTasks services method
func (ts *TaskHandler) GetTasks(ctx *gin.Context) {
	tasks, err := ts.TaskService.GetAllTasks()
	if err != nil {
		slog.Error("An error occured while getting tasks", "Error", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError())
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

// Add new task
// Uses the Create Task service

func (ts *TaskHandler) HandleCreate(ctx *gin.Context) {
	var params dto.TaskCreateRequest

	// check if request body checks with request struct
	if err := ctx.ShouldBindJSON(&params); err != nil {
		slog.Error("Invalid body request", "Error", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequest())
		return
	}

	// create new task
	newTask := data.Task{
		Title: params.Title,
	}

	// add task to database
	task, err := ts.TaskService.CreateNewTask(newTask)
	if err != nil {
		slog.Error("Error occured while saving task", "Error", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError())
		return
	}

	// return response
	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"task":    task,
		"message": "task created successfully",
	})

}
