package useCases

import (
	"github.com/jansen44/nagasu-kujira/core/entities"
	"github.com/jansen44/nagasu-kujira/core/repositories"
)

type ITasksUseCases interface {
	AddNewTask(name string, description string, missionID int) (*entities.TasksEntity, error)
	UpdateOneTask(name, description string, ID int64) (*entities.TasksEntity, error)
	RemoveOneTask(ID int64) (*entities.TasksEntity, error)
}

type TasksUseCases struct {
	tasksRepository repositories.ITasksRepository
}

func NewTasksUseCases(tasksRepository repositories.ITasksRepository) ITasksUseCases {
	return &TasksUseCases{
		tasksRepository,
	}
}

func (t TasksUseCases) AddNewTask(name string, description string, missionID int) (*entities.TasksEntity, error) {
	return t.tasksRepository.CreateTask(&entities.TasksEntity{Name: name, Description: description, MissionID: missionID})
}

func (t TasksUseCases) UpdateOneTask(name, description string, ID int64) (*entities.TasksEntity, error) {
	return t.tasksRepository.UpdateTask(&entities.TasksEntity{ID: ID, Name: name, Description: description})
}

func (t TasksUseCases) RemoveOneTask(ID int64) (*entities.TasksEntity, error) {
	return t.tasksRepository.DeleteTask(ID)
}
