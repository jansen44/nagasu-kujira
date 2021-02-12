package repositories

import (
	"context"
	"github.com/jansen44/nagasu-kujira/core/entities"
)

type ITasksRepository interface {
	CreateTask(ctx context.Context, task *entities.TasksEntity) (*entities.TasksEntity, error)
	UpdateTask(ctx context.Context, task *entities.TasksEntity) (*entities.TasksEntity, error)
	DeleteTask(ctx context.Context, ID int64) (*entities.TasksEntity, error)
}
