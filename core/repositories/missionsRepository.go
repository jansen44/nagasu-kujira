package repositories

import "github.com/jansen44/nagasu-kujira/core/entities"

type IMissionsRepository interface {
	// CRUD ==========================================================
	CreateMission(mission *entities.MissionsEntity, projectID int) (*entities.MissionsEntity, error)
	ReadMissions(projectID int) ([]entities.MissionsEntity, error)
	UpdateMission(mission *entities.MissionsEntity) (*entities.MissionsEntity, error)
	DeleteMission(ID int) (*entities.MissionsEntity, error)
}
