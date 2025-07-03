package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type Task struct {
	ID        *uint      `gorm:"primaryKey" json:"id"`
	Title     *string    `json:"title"`
	Content   *string    `json:"content"`
	Deadline  *time.Time `json:"deadline"`
	Done      *bool      `json:"done"`
	CreatedAt *time.Time `json:"created_at"`
}

func (Task) TableName() string { return "tasks" }

func main() {
	var task Task
	dsn := "root:@tcp(localhost:3306)/practice?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&task); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	r := gin.Default()
	err = r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("failed to set trusted proxies: %v", err)
	}
	r.POST("/tasks", func(c *gin.Context) {
		var tasks Task
		if err := c.BindJSON(&tasks); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&tasks)
		fmt.Println("Created task: ", tasks)
		c.JSON(200, tasks)
	})
	r.GET("/tasks", func(c *gin.Context) {
		var tasks []Task
		db.Find(&tasks)
		c.JSON(200, tasks)
	})
	r.GET("/tasks/overdue", func(c *gin.Context) {
		var tasks []Task
		db.Where("done = ? AND deadline < ?", false, time.Now()).Find(&tasks)
		c.JSON(200, tasks)
	})
	r.GET("/tasks/:id", func(c *gin.Context) {
		var task Task
		id := c.Param("id")
		db.First(&task, id)
		c.JSON(200, task)
	})
	r.PUT("/tasks/:id", func(c *gin.Context) {
		var task Task
		id := c.Param("id")
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		updates := make(map[string]interface{})
		if task.Title != nil {
			updates["title"] = task.Title
		}
		if task.Content != nil {
			updates["content"] = task.Content
		}
		if task.Deadline != nil {
			updates["deadline"] = task.Deadline
		}
		if task.Done != nil {
			updates["done"] = task.Done
		}
		db.Model(&task).Where("id = ?", id).Updates(updates)
		c.JSON(200, task)
	})
	r.DELETE("/tasks/:id", func(c *gin.Context) {
		var task Task
		id := c.Param("id")
		db.First(&task, id).Delete(&task)
		c.JSON(200, task)
	})
	r.GET("/tasks/search", func(c *gin.Context) {
		var tasks []Task
		title := "%" + c.Query("title") + "%"
		db.Where("title LIKE ?", title).Find(&tasks)
		c.JSON(200, tasks)
	})
	r.GET("/tasks/today", func(c *gin.Context) {
		var tasks []Task
		db.Where("DATE(deadline) = ?", time.Now().Format("2006-01-02")).Find(&tasks)
		c.JSON(200, tasks)
	})
	r.Run(":8080")
}
