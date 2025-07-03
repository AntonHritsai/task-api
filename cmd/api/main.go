package main

import (
	"github.com/AntonKhPI2/task-api/internal/database"
	"github.com/AntonKhPI2/task-api/internal/handlers"
	"github.com/AntonKhPI2/task-api/internal/repositories"
	"github.com/AntonKhPI2/task-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	db := database.InitDB()
	taskRepository := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)
	r := gin.Default()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("failed to set trusted proxies: %v", err)
	}
	r.POST("/tasks", taskHandler.PostTask)
	r.GET("/tasks", taskHandler.GetAllTasks)
	r.GET("/tasks/overdue", taskHandler.GetTaskOverDue)
	r.GET("/tasks/:id", taskHandler.GetTaskByID)
	r.PUT("/tasks/:id", taskHandler.ChangeTaskByID)
	r.DELETE("/tasks/:id", taskHandler.DeleteTaskByID)
	r.GET("/tasks/search", taskHandler.FindTasksByTitle)
	r.GET("/tasks/today", taskHandler.GetTasksForToday)
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
