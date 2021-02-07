package mysql

import (
	"context"
	"database/sql"
	"github.com/jansen44/nagasu-kujira/core/entities"
	"github.com/jansen44/nagasu-kujira/core/repositories"
	"github.com/jansen44/nagasu-kujira/infra/mysql/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type MissionsMySQL struct {
	db  *sql.DB
	ctx context.Context
}

// Util ===============================================================================
func NewMissionsMySQL(db *sql.DB, ctx context.Context) repositories.IMissionsRepository {
	return &MissionsMySQL{db, ctx}
}

func missionModelToEntity(mission *models.Mission) *entities.MissionsEntity {
	return entities.NewMissionsEntity(mission.Name, mission.ProjectsID, mission.ID)
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
func (m MissionsMySQL) CreateMission(mission *entities.MissionsEntity) (*entities.MissionsEntity, error) {
	model := models.Mission{Name: mission.Name, ProjectsID: mission.ProjectID}
	err := model.Insert(m.ctx, m.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return missionModelToEntity(&model), nil
}

func (m MissionsMySQL) UpdateMission(mission *entities.MissionsEntity) (*entities.MissionsEntity, error) {
	model, err := models.FindMission(m.ctx, m.db, mission.ID)
	if err != nil {
		return nil, err
	}
	missionEntityToModel(mission, model)
	_, err = model.Update(m.ctx, m.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return missionModelToEntity(model), nil
}

func (m MissionsMySQL) DeleteMission(ID int) (*entities.MissionsEntity, error) {
	model, err := models.FindMission(m.ctx, m.db, ID)
	if err != nil {
		return nil, err
	}
	_, err = model.Delete(m.ctx, m.db)
	if err != nil {
		return nil, err
	}
	return missionModelToEntity(model), nil
}
