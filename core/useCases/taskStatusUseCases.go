package useCases

import (
	"context"
	"github.com/jansen44/nagasu-kujira/core/entities"
	"github.com/jansen44/nagasu-kujira/core/repositories"
)

type ITasksStatusUseCases interface {
	AddNewTaskStatus(ctx context.Context, name string, order int, missionID int) (*entities.TaskStatusEntity, error)
	UpdateOneTaskStatus(ctx context.Context, name string, order int, ID int64) (*entities.TaskStatusEntity, error)
	RemoveOneTaskStatus(ctx context.Context, ID int64) (*entities.TaskStatusEntity, error)
}

type TaskStatusUseCases struct {
	taskStatusRepository repositories.ITaskStatusRepository
}

func NewTaskStatusUseCases(taskStatusRepository repositories.ITaskStatusRepository) ITasksStatusUseCases {
	return &TaskStatusUseCases{
		taskStatusRepository,
	}
}

func (ts TaskStatusUseCases) AddNewTaskStatus(ctx context.Context, name string, order int, missionID int) (*entities.TaskStatusEntity, error) {
	return ts.taskStatusRepository.CreateTaskStatus(ctx, &entities.TaskStatusEntity{Name: name, Order: order, MissionID: missionID})
}

func (ts TaskStatusUseCases) UpdateOneTaskStatus(ctx context.Context, name string, order int, ID int64) (*entities.TaskStatusEntity, error) {
	return ts.taskStatusRepository.UpdateTaskStatus(ctx, &entities.TaskStatusEntity{ID: ID, Name: name, Order: order})
}

func (ts TaskStatusUseCases) RemoveOneTaskStatus(ctx context.Context, ID int64) (*entities.TaskStatusEntity, error) {
	return ts.taskStatusRepository.DeleteTaskStatus(ctx, ID)
}
