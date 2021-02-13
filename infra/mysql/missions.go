package mysql

import (
	"context"
	"database/sql"
	"github.com/jansen44/nagasu-kujira/core/entities"
	"github.com/jansen44/nagasu-kujira/core/repositories"
	"github.com/jansen44/nagasu-kujira/infra/mysql/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MissionsMySQL struct {
	db *sql.DB
}

// Util ===============================================================================
func NewMissionsMySQL(db *sql.DB) repositories.IMissionsRepository {
	return &MissionsMySQL{db}
}

func missionModelToEntity(model *models.Mission) *entities.MissionsEntity {
	mission := entities.NewMissionsEntity(model.Name, model.ProjectsID, model.ID)
	if model.R != nil {
		if model.R.Tasks != nil {
			mission.Tasks = taskModelSliceToPointerEntity(model.R.Tasks)
		}
		if model.R.TaskStatuses != nil {
			mission.TaskStatus = taskStatusModelSliceToPointerEntity(model.R.TaskStatuses)
		}
	}
	return mission
}

func missionModelSliceToEntity(missionSlice models.MissionSlice) []entities.MissionsEntity {
	missions := make([]entities.MissionsEntity, 0)
	for _, mission := range missionSlice {
		missions = append(missions, *missionModelToEntity(mission))
	}
	return missions
}

func missionModelSliceToPointerEntity(missionSlice models.MissionSlice) []*entities.MissionsEntity {
	missions := make([]*entities.MissionsEntity, 0)
	for _, mission := range missionSlice {
		missions = append(missions, missionModelToEntity(mission))
	}
	return missions
}

func missionEntityToModel(entity *entities.MissionsEntity, model *models.Mission) {
	model.ID = entity.ID
	model.Name = entity.Name
}

// Repository Methods ==================================================================
func (m MissionsMySQL) CreateMission(ctx context.Context, mission *entities.MissionsEntity) (*entities.MissionsEntity, error) {
	model := models.Mission{Name: mission.Name, ProjectsID: mission.ProjectID}
	err := model.Insert(ctx, m.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return missionModelToEntity(&model), nil
}

func (m MissionsMySQL) UpdateMission(ctx context.Context, mission *entities.MissionsEntity) (*entities.MissionsEntity, error) {
	model, err := models.FindMission(ctx, m.db, mission.ID)
	if err != nil {
		return nil, err
	}
	missionEntityToModel(mission, model)
	_, err = model.Update(ctx, m.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return missionModelToEntity(model), nil
}

func (m MissionsMySQL) DeleteMission(ctx context.Context, ID int) (*entities.MissionsEntity, error) {
	model, err := models.FindMission(ctx, m.db, ID)
	if err != nil {
		return nil, err
	}
	_, err = model.Delete(ctx, m.db)
	if err != nil {
		return nil, err
	}
	return missionModelToEntity(model), nil
}

func (m MissionsMySQL) ReadMission(ctx context.Context, ID int) (*entities.MissionsEntity, error) {
	model, err := models.Missions(
		qm.Load(models.MissionRels.Tasks),
		qm.Load(models.MissionRels.TaskStatuses),
		qm.Where("id=?", ID),
	).One(ctx, m.db)
	if err != nil {
		return nil, err
	}
	return missionModelToEntity(model), nil
}
