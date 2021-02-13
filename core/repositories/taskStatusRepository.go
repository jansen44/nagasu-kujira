package repositories

import (
	"context"
	"github.com/jansen44/nagasu-kujira/core/entities"
)

type ITaskStatusRepository interface {
	CreateTaskStatus(ctx context.Context, taskStatus *entities.TaskStatusEntity) (*entities.TaskStatusEntity, error)
	UpdateTaskStatus(ctx context.Context, taskStatus *entities.TaskStatusEntity) (*entities.TaskStatusEntity, error)
	DeleteTaskStatus(ctx context.Context, ID int64) (*entities.TaskStatusEntity, error)
}
