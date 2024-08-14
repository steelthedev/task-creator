package services

import (
	"github.com/steelthedev/task-handler/data"
	"github.com/steelthedev/task-handler/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func (ts *TaskService) CreateNewTask(newTask data.Task) (*data.Task, error) {

	task, err := ts.repo.Create(newTask)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (ts *TaskService) GetAllTasks() ([]*data.Task, error) {

	tasks, err := ts.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ts *TaskService) GetTask(id uint) (*data.Task, error) {

	task, err := ts.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}
