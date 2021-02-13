package mysql

import (
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"
	"github.com/jansen44/nagasu-kujira/core/entities"
	"github.com/jansen44/nagasu-kujira/core/repositories"
	"github.com/jansen44/nagasu-kujira/infra/mysql/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TasksMySQL struct {
	db *sql.DB
}

// Util ===============================================================================
func NewTasksMySQL(db *sql.DB) repositories.ITasksRepository {
	return &TasksMySQL{db}
}

func taskModelToEntity(model *models.Task) *entities.TasksEntity {
	task := entities.NewTasksEntity(model.Name, model.Description, model.MissionID, model.CurrentStatus, model.ID)
	return task
}

func taskModelSliceToPointerEntity(taskSlice models.TaskSlice) []*entities.TasksEntity {
	tasks := make([]*entities.TasksEntity, 0)
	for _, task := range taskSlice {
		tasks = append(tasks, taskModelToEntity(task))
	}
	return tasks
}

func taskEntityToModel(entity *entities.TasksEntity, model *models.Task) {
	model.ID = entity.ID
	model.Name = entity.Name
	model.Description = entity.Description
	model.CurrentStatus = entity.TaskStatusID
}

// Repository Methods ==================================================================
func (t TasksMySQL) CreateTask(ctx context.Context, task *entities.TasksEntity) (*entities.TasksEntity, error) {
	// Before inserting a task, there's a need in validating if the status is owned by the task mission.
	taskStatus, err := models.FindTaskStatus(ctx, t.db, task.TaskStatusID)
	if err != nil {
		return nil, err
	}
	if taskStatus.MissionID != task.MissionID {
		return nil, errors.New("invalid status")
	}
	model := models.Task{
		Name:          task.Name,
		Description:   task.Description,
		MissionID:     task.MissionID,
		CurrentStatus: task.TaskStatusID,
	}
	err = model.Insert(ctx, t.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return taskModelToEntity(&model), nil
}

func (t TasksMySQL) UpdateTask(ctx context.Context, task *entities.TasksEntity) (*entities.TasksEntity, error) {
	model, err := models.FindTask(ctx, t.db, task.ID)
	if err != nil {
		return nil, err
	}
	// Before updating a task, there's a need in validating if the status is owned by the task mission
	// (only if it's different from the current model)
	if model.CurrentStatus != task.TaskStatusID {
		taskStatus, err := models.FindTaskStatus(ctx, t.db, task.TaskStatusID)
		if err != nil {
			return nil, err
		}
		if taskStatus.MissionID != model.MissionID {
			return nil, errors.New("invalid status")
		}
	}
	taskEntityToModel(task, model)
	_, err = model.Update(ctx, t.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return taskModelToEntity(model), nil
}

func (t TasksMySQL) DeleteTask(ctx context.Context, ID int64) (*entities.TasksEntity, error) {
	model, err := models.FindTask(ctx, t.db, ID)
	if err != nil {
		return nil, err
	}
	_, err = model.Delete(ctx, t.db)
	if err != nil {
		return nil, err
	}
	return taskModelToEntity(model), nil
}
