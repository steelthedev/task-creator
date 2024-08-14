package main

import (
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/joho/godotenv"
	"github.com/steelthedev/task-handler/conn"
	"github.com/steelthedev/task-handler/data"
	"github.com/steelthedev/task-handler/services"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// Set logger file dist
	log.SetOutput(&lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

	// Initiate environment.
	if err := godotenv.Load("./.env"); err != nil {
		slog.Info("Error connecting to dotenv")
	}

	// Start Database
	if err := conn.InitDB(); err != nil {
		slog.Error("Could not connect to db", "Error=", err)
	}

	s := services.TaskService{}

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

	_, err := s.CreateNewTask(tasktwo)
	if err != nil {
		slog.Info("could not add a new task", "Error=", err)
		return
	}
	_, err = s.CreateNewTask(taskThree)
	if err != nil {
		slog.Info("could not add a new task", "Error=", err)
		return
	}

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

	fmt.Println("Single mf")
	task, err := s.GetTask(2)
	if err != nil {
		return
	}
	fmt.Println(task)
}
