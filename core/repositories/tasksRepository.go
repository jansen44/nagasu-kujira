package repositories

import "github.com/jansen44/nagasu-kujira/core/entities"

type ITasksRepository interface {
	CreateTask(task *entities.TasksEntity) (*entities.TasksEntity, error)
	UpdateTask(task *entities.TasksEntity) (*entities.TasksEntity, error)
	DeleteTask(ID int64) (*entities.TasksEntity, error)
}
