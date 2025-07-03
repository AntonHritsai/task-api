package repositories

import (
	"context"
	"github.com/AntonKhPI2/task-api/internal/models"
	"gorm.io/gorm"
	"time"
)

type TaskRepository interface {
	PostTask(ctx context.Context, repository *models.Task) error
	GetAllTasks(c context.Context) ([]models.Task, error)
	GetTaskOverdueByDate(c context.Context, date time.Time) ([]models.Task, error)
	GetTaskByID(c context.Context, taskID string) (*models.Task, error)
	ChangeTaskByID(c context.Context, taskID uint64, newTask *models.Task) (*models.Task, error)
	DeleteTaskByID(c context.Context, taskID uint64) error
	FindTasksByTitle(c context.Context, title string) ([]models.Task, error)
	FindTasksByDate(c context.Context, date time.Time) ([]models.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) PostTask(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *taskRepository) GetAllTasks(c context.Context) ([]models.Task, error) {
	var tasks []models.Task
	db := r.db.WithContext(c)
	if err := db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) GetTaskOverdueByDate(c context.Context, date time.Time) ([]models.Task, error) {
	var tasks []models.Task
	db := r.db.WithContext(c)
	if err := db.Where("done = ? AND deadline < ?", false, date).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) GetTaskByID(c context.Context, taskID string) (*models.Task, error) {
	var task models.Task
	db := r.db.WithContext(c)
	if err := db.First(&task, taskID).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) ChangeTaskByID(c context.Context, taskID uint64, task *models.Task) (*models.Task, error) {
	result := r.db.WithContext(c).Model(&models.Task{}).Updates(task)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	var updatedTask models.Task
	if err := r.db.First(&updatedTask, taskID).Error; err != nil {
		return nil, err
	}
	return &updatedTask, nil
}

func (r *taskRepository) DeleteTaskByID(c context.Context, taskID uint64) error {
	result := r.db.WithContext(c).Where("id = ?", taskID).Delete(&models.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *taskRepository) FindTasksByTitle(c context.Context, title string) ([]models.Task, error) {
	var tasks []models.Task
	db := r.db.WithContext(c)
	if err := db.Where("title LIKE ?", title).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) FindTasksByDate(c context.Context, date time.Time) ([]models.Task, error) {
	var tasks []models.Task
	db := r.db.WithContext(c)
	if err := db.Where("DATE(deadline) = ?", date.Format("2006-01-02")).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
