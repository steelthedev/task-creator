package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/steelthedev/task-handler/conn"

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

	// s := services.TaskService{}

	app := gin.Default()

	taskRoutes := app.Group("task")

	// Get port from env
	port := os.Getenv("PORT")

	log.Fatal(app.Run(port))

}
