package main

import (
	"fmt"
	"time"

	"github.com/steelthedev/task-handler/data"
)

func main() {

	s := data.TaskService{}

	taskOne := data.Task{
		Title:     "Task 1",
		CreatedAt: time.Now(),
		EndAt:     time.Now(),
	}

	tasktwo := data.Task{
		Title:     "Task 2",
		CreatedAt: time.Now(),
		EndAt:     time.Now(),
	}

	taskThree := data.Task{
		Title:     "Task 3",
		CreatedAt: time.Now(),
		EndAt:     time.Now(),
	}

	s.CreateNewTask(taskOne)
	s.CreateNewTask(tasktwo)
	s.CreateNewTask(taskThree)

	tasks, err := s.GetAllTasks()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Total Tasks: %d /n", len(tasks))

	for index, task := range tasks {
		fmt.Println("-----------------------------------")
		fmt.Printf(" %d. %s \n Start time: %s \n End time: %s \n", index, task.Title, task.CreatedAt, task.EndAt)
		fmt.Println("------------------------------------")
	}
}
