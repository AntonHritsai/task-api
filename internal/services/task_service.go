package services

import (
	"context"
	"errors"
	"github.com/AntonKhPI2/task-api/internal/models"
	"github.com/AntonKhPI2/task-api/internal/repositories"
	"strconv"
	"time"
)

type TaskService interface {
	PostTask(ctx context.Context, task models.TaskRequest) (*models.Task, error)
	GetAllTasks(ctx context.Context) ([]models.Task, error)
	GetTaskOverdue(ctx context.Context) ([]models.Task, error)
	GetTaskByID(ctx context.Context, id string) (*models.Task, error)
	ChangeTaskByID(ctx context.Context, id string, task *models.TaskUpdateRequest) (*models.Task, error)
	DeleteTaskByID(ctx context.Context, id string) error
	FindTasksByTitle(c context.Context, title string) ([]models.Task, error)
	GetTasksForToday(ctx context.Context) ([]models.Task, error)
}
type taskService struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) PostTask(ctx context.Context, task models.TaskRequest) (*models.Task, error) {
	doneStatus := false
	newTask := &models.Task{
		Title:    &task.Title,
		Content:  &task.Content,
		Deadline: task.Deadline,
		Done:     &doneStatus,
	}
	err := s.repo.PostTask(ctx, newTask)
	if err != nil {
		return nil, err
	}
	return newTask, nil
}
func (s *taskService) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	tasks, err := s.repo.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *taskService) GetTaskOverdue(ctx context.Context) ([]models.Task, error) {
	tasks, err := s.repo.GetTaskOverdueByDate(ctx, time.Now())
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *taskService) GetTaskByID(ctx context.Context, id string) (*models.Task, error) {
	return s.repo.GetTaskByID(ctx, id)
}

func (s *taskService) ChangeTaskByID(ctx context.Context, id string, task *models.TaskUpdateRequest) (*models.Task, error) {
	taskID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	taskToUpdate := &models.Task{
		Title:    task.Title,
		Content:  task.Content,
		Deadline: task.Deadline,
		Done:     task.Done,
	}

	return s.repo.ChangeTaskByID(ctx, taskID, taskToUpdate)
}

func (s *taskService) DeleteTaskByID(ctx context.Context, id string) error {
	taskID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("invalid task ID")
	}
	err = s.repo.DeleteTaskByID(ctx, taskID)
	if err != nil {
		return err
	}
	return nil
}

func (s *taskService) FindTasksByTitle(c context.Context, title string) ([]models.Task, error) {
	return s.repo.FindTasksByTitle(c, "%"+title+"%")
}

func (s *taskService) GetTasksForToday(ctx context.Context) ([]models.Task, error) {
	return s.repo.FindTasksByDate(ctx, time.Now())
}
