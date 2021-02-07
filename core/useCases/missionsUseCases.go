package useCases

import (
	"github.com/jansen44/nagasu-kujira/core/entities"
	"github.com/jansen44/nagasu-kujira/core/repositories"
)

type IMissionsUseCases interface {
	AddNewMission(name string) (*entities.MissionsEntity, error)
	UpdateOneMission(name string, ID int) (*entities.MissionsEntity, error)
	RemoveOneMission(ID int) (*entities.MissionsEntity, error)
}

type MissionsUseCases struct {
	missionsRepository repositories.IMissionsRepository
}

func NewMissionsUseCases(missionsRepository repositories.IMissionsRepository) IMissionsUseCases {
	return &MissionsUseCases{
		missionsRepository: missionsRepository,
	}
}

func (m MissionsUseCases) AddNewMission(name string) (*entities.MissionsEntity, error) {
	return m.missionsRepository.CreateMission(&entities.MissionsEntity{Name: name})
}

func (m MissionsUseCases) UpdateOneMission(name string, ID int) (*entities.MissionsEntity, error) {
	return m.missionsRepository.UpdateMission(&entities.MissionsEntity{ID: ID, Name: name})
}

func (m MissionsUseCases) RemoveOneMission(ID int) (*entities.MissionsEntity, error) {
	return m.missionsRepository.DeleteMission(ID)
}