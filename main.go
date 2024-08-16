package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/steelthedev/task-handler/conn"
	"github.com/steelthedev/task-handler/handlers"
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

	taskHandler := handlers.NewTaskHandler(s)

	app := gin.Default()

	// Register task route groups and handlers
	taskRoutes := app.Group("task")
	taskRoutes.GET("/all", taskHandler.GetTasks)
	taskRoutes.POST("/add", taskHandler.HandleCreate)

	// Get port from env
	port := os.Getenv("PORT")
	if len(port) > 0 {
		log.Fatal(app.Run(port))
	} else {
		log.Fatal(app.Run(":8000"))
	}

}
