package mysql

import (
	"context"
	"database/sql"
	"github.com/jansen44/nagasu-kujira/core/entities"
	"github.com/jansen44/nagasu-kujira/core/repositories"
	"github.com/jansen44/nagasu-kujira/infra/mysql/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TasksMySQL struct {
	db  *sql.DB
	ctx context.Context
}

// Util ===============================================================================
func NewTasksMySQL(db *sql.DB, ctx context.Context) repositories.ITasksRepository {
	return &TasksMySQL{db, ctx}
}

func taskModelToEntity(model *models.Task) *entities.TasksEntity {
	task := entities.NewTasksEntity(model.Name, model.Description, model.MissionID, model.ID)
	return task
}

func taskEntityToModel(entity *entities.TasksEntity, model *models.Task) {
	model.ID = entity.ID
	model.Name = entity.Name
	model.Description = entity.Description
}

// Repository Methods ==================================================================
func (t TasksMySQL) CreateTask(task *entities.TasksEntity) (*entities.TasksEntity, error) {
	model := models.Task{Name: task.Name, Description: task.Description, MissionID: task.MissionID}
	err := model.Insert(t.ctx, t.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return taskModelToEntity(&model), nil
}

func (t TasksMySQL) UpdateTask(task *entities.TasksEntity) (*entities.TasksEntity, error) {
	model, err := models.FindTask(t.ctx, t.db, task.ID)
	if err != nil {
		return nil, err
	}
	taskEntityToModel(task, model)
	_, err = model.Update(t.ctx, t.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return taskModelToEntity(model), nil
}

func (t TasksMySQL) DeleteTask(ID int64) (*entities.TasksEntity, error) {
	model, err := models.FindTask(t.ctx, t.db, ID)
	if err != nil {
		return nil, err
	}
	_, err = model.Delete(t.ctx, t.db)
	if err != nil {
		return nil, err
	}
	return taskModelToEntity(model), nil
}
