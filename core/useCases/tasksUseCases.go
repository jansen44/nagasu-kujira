package useCases

import (
	"context"
	"github.com/jansen44/nagasu-kujira/core/entities"
	"github.com/jansen44/nagasu-kujira/core/repositories"
)

type ITasksUseCases interface {
	AddNewTask(ctx context.Context, name string, description string, missionID int, statusID int64) (*entities.TasksEntity, error)
	UpdateOneTask(ctx context.Context, name, description string, statusID, ID int64) (*entities.TasksEntity, error)
	RemoveOneTask(ctx context.Context, ID int64) (*entities.TasksEntity, error)
}

type TasksUseCases struct {
	tasksRepository repositories.ITasksRepository
}

func NewTasksUseCases(tasksRepository repositories.ITasksRepository) ITasksUseCases {
	return &TasksUseCases{
		tasksRepository,
	}
}

func (t TasksUseCases) AddNewTask(ctx context.Context, name string, description string, missionID int, statusID int64) (*entities.TasksEntity, error) {
	return t.tasksRepository.CreateTask(ctx, &entities.TasksEntity{
		Name:         name,
		Description:  description,
		MissionID:    missionID,
		TaskStatusID: statusID,
	})
}

func (t TasksUseCases) UpdateOneTask(ctx context.Context, name, description string, statusID, ID int64) (*entities.TasksEntity, error) {
	return t.tasksRepository.UpdateTask(ctx, &entities.TasksEntity{
		ID:           ID,
		Name:         name,
		Description:  description,
		TaskStatusID: statusID,
	})
}

func (t TasksUseCases) RemoveOneTask(ctx context.Context, ID int64) (*entities.TasksEntity, error) {
	return t.tasksRepository.DeleteTask(ctx, ID)
}
