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

//
//func projectModelSliceToEntity(projectSlice models.ProjectSlice) []entities.ProjectsEntity {
//	projects := make([]entities.ProjectsEntity, 0)
//	for _, project := range projectSlice {
//		projects = append(projects, *projectModelToEntity(project))
//	}
//	return projects
//}
//
//func projectEntityToModel(entity *entities.ProjectsEntity, model *models.Project) {
//	model.ID = entity.ID
//	model.Name = entity.Name
//}

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
	panic("implement me")
}

func (m MissionsMySQL) DeleteMission(ID int) (*entities.MissionsEntity, error) {
	panic("implement me")
}
