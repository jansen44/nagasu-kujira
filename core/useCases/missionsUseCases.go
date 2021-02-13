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
	taskStatusUseCases ITasksStatusUseCases
	missionsRepository repositories.IMissionsRepository
}

func NewMissionsUseCases(missionsRepository repositories.IMissionsRepository, taskStatusUseCases ITasksStatusUseCases) IMissionsUseCases {
	return &MissionsUseCases{
		missionsRepository: missionsRepository,
		taskStatusUseCases: taskStatusUseCases,
	}
}

func (m MissionsUseCases) AddNewMission(ctx context.Context, name string, projectId int) (*entities.MissionsEntity, error) {
	mission, err := m.missionsRepository.CreateMission(ctx, &entities.MissionsEntity{Name: name, ProjectID: projectId})
	if err != nil {
		return nil, err
	}
	// By default, every mission needs to be initialized with default Task Status.
	for i, taskStatusName := range entities.GetDefaultTaskStatuses() {
		status, err := m.taskStatusUseCases.AddNewTaskStatus(ctx, string(taskStatusName), i, mission.ID)
		if err != nil {
			return nil, err
		}
		mission.TaskStatus = append(mission.TaskStatus, status)
	}
	return mission, nil
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
