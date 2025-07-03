package handlers

import (
	"github.com/AntonKhPI2/task-api/internal/models"
	"github.com/AntonKhPI2/task-api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskHandler interface {
	PostTask(c *gin.Context)
	GetAllTasks(c *gin.Context)
	GetTaskOverDue(c *gin.Context)
	GetTaskByID(c *gin.Context)
	ChangeTaskByID(c *gin.Context)
	DeleteTaskByID(c *gin.Context)
	FindTasksByTitle(c *gin.Context)
	GetTasksForToday(c *gin.Context)
}

type taskHandler struct {
	svc services.TaskService
}

func NewTaskHandler(svc services.TaskService) TaskHandler {
	return &taskHandler{svc: svc}
}

func (h *taskHandler) PostTask(c *gin.Context) {
	var task models.TaskRequest
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTask, err := h.svc.PostTask(c.Request.Context(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusCreated, createdTask)
}

func (h *taskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.svc.GetAllTasks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *taskHandler) GetTaskOverDue(c *gin.Context) {
	tasks, err := h.svc.GetTaskOverdue(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *taskHandler) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := h.svc.GetTaskByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *taskHandler) ChangeTaskByID(c *gin.Context) {
	var task models.TaskUpdateRequest
	id := c.Param("id")
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTask, err := h.svc.ChangeTaskByID(c.Request.Context(), id, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (h *taskHandler) DeleteTaskByID(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.DeleteTaskByID(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *taskHandler) FindTasksByTitle(c *gin.Context) {
	tasks, err := h.svc.FindTasksByTitle(c.Request.Context(), c.Query("title"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *taskHandler) GetTasksForToday(c *gin.Context) {
	tasks, err := h.svc.GetTasksForToday(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
