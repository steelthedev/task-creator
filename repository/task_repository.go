package repository

import (
	"github.com/steelthedev/task-handler/conn"
	"github.com/steelthedev/task-handler/data"
)

type taskInterface interface {
	Create(task *data.Task) (*data.Task, error)
	GetAll() (*[]data.Task, error)
	Get(id uint) (*data.Task, error)
}

type TaskRepository struct{}

func (t *TaskRepository) Create(task data.Task) (*data.Task, error) {

	if err := conn.DB.Create(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *TaskRepository) GetAll() ([]*data.Task, error) {
	var tasks []*data.Task

	if err := conn.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *TaskRepository) Get(id uint) (*data.Task, error) {
	var task data.Task

	if err := conn.DB.Where("ID=?", id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
