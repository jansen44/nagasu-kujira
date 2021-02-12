package useCases

import (
	"context"
	"github.com/jansen44/nagasu-kujira/core/entities"
	"github.com/jansen44/nagasu-kujira/core/repositories"
)

type IMissionsUseCases interface {
	AddNewMission(ctx context.Context, name string, projectId int) (*entities.MissionsEntity, error)
	UpdateOneMission(ctx context.Context, name string, ID int) (*entities.MissionsEntity, error)
	RemoveOneMission(ctx context.Context, ID int) (*entities.MissionsEntity, error)
	GetMissionInfo(ctx context.Context, ID int) (*entities.MissionsEntity, error)
}

type MissionsUseCases struct {
	missionsRepository repositories.IMissionsRepository
}

func NewMissionsUseCases(missionsRepository repositories.IMissionsRepository) IMissionsUseCases {
	return &MissionsUseCases{
		missionsRepository: missionsRepository,
	}
}

func (m MissionsUseCases) AddNewMission(ctx context.Context, name string, projectId int) (*entities.MissionsEntity, error) {
	return m.missionsRepository.CreateMission(ctx, &entities.MissionsEntity{Name: name, ProjectID: projectId})
}

func (m MissionsUseCases) UpdateOneMission(ctx context.Context, name string, ID int) (*entities.MissionsEntity, error) {
	return m.missionsRepository.UpdateMission(ctx, &entities.MissionsEntity{ID: ID, Name: name})
}

func (m MissionsUseCases) RemoveOneMission(ctx context.Context, ID int) (*entities.MissionsEntity, error) {
	return m.missionsRepository.DeleteMission(ctx, ID)
}

func (m MissionsUseCases) GetMissionInfo(ctx context.Context, ID int) (*entities.MissionsEntity, error) {
	return m.missionsRepository.ReadMission(ctx, ID)
}
