package data

import (
	"fmt"
	"sync"
	"time"
)

type taskInterface interface {
	CreateTask(task *Task) (*Task, error)
	GetAll() (*[]Task, error)
}

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"task" gorm:"column:title"`
	CreatedAt time.Time `json:"created_at"`
	EndAt     time.Time `json:"end_at"`
}

type Repository struct {
	mu    sync.RWMutex
	tasks []*Task
}

func (r *Repository) CreateTask(task Task) (*Task, error) {
	r.mu.Lock()
	update := append(r.tasks, &task)
	r.tasks = update
	r.mu.Unlock()
	return &task, nil
}

func (r *Repository) GetAll() ([]*Task, error) {
	if len(r.tasks) > 0 {
		return r.tasks, nil
	}
	return nil, fmt.Errorf("No task found")
}

type TaskService struct {
	repo Repository
}

func (ts *TaskService) CreateNewTask(newTask Task) (*Task, error) {

	task, err := ts.repo.CreateTask(newTask)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (ts *TaskService) GetAllTasks() ([]*Task, error) {

	tasks, err := ts.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
