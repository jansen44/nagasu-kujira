package mysql

import (
	"context"
	"database/sql"
	"github.com/jansen44/nagasu-kujira/core/entities"
	"github.com/jansen44/nagasu-kujira/core/repositories"
	"github.com/jansen44/nagasu-kujira/infra/mysql/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TaskStatusMySQL struct {
	db *sql.DB
}

// Util ===============================================================================
func NewTaskStatusMySQL(db *sql.DB) repositories.ITaskStatusRepository {
	return &TaskStatusMySQL{db}
}

func taskStatusModelToEntity(model *models.TaskStatus) *entities.TaskStatusEntity {
	taskStatus := entities.NewTaskStatusEntity(model.Name, model.MissionID, model.StatusOrder, model.ID)
	return taskStatus
}

func taskStatusModelSliceToPointerEntity(taskStatusSlice models.TaskStatusSlice) []*entities.TaskStatusEntity {
	taskStatuses := make([]*entities.TaskStatusEntity, 0)
	for _, taskStatus := range taskStatusSlice {
		taskStatuses = append(taskStatuses, taskStatusModelToEntity(taskStatus))
	}
	return taskStatuses
}

func taskStatusEntityToModel(entity *entities.TaskStatusEntity, model *models.TaskStatus) {
	model.ID = entity.ID
	model.Name = entity.Name
	model.StatusOrder = entity.Order
}

// Repository Methods ==================================================================
func (ts TaskStatusMySQL) CreateTaskStatus(ctx context.Context, taskStatus *entities.TaskStatusEntity) (*entities.TaskStatusEntity, error) {
	model := models.TaskStatus{Name: taskStatus.Name, StatusOrder: taskStatus.Order, MissionID: taskStatus.MissionID}
	err := model.Insert(ctx, ts.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return taskStatusModelToEntity(&model), nil
}

func (ts TaskStatusMySQL) UpdateTaskStatus(ctx context.Context, taskStatus *entities.TaskStatusEntity) (*entities.TaskStatusEntity, error) {
	model, err := models.FindTaskStatus(ctx, ts.db, taskStatus.ID)
	if err != nil {
		return nil, err
	}
	taskStatusEntityToModel(taskStatus, model)
	_, err = model.Update(ctx, ts.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return taskStatusModelToEntity(model), nil
}

func (ts TaskStatusMySQL) DeleteTaskStatus(ctx context.Context, ID int64) (*entities.TaskStatusEntity, error) {
	model, err := models.FindTaskStatus(ctx, ts.db, ID)
	if err != nil {
		return nil, err
	}
	_, err = model.Delete(ctx, ts.db)
	if err != nil {
		return nil, err
	}
	return taskStatusModelToEntity(model), nil
}
