package repositories

import (
	"context"
	"github.com/jansen44/nagasu-kujira/core/entities"
)

type IMissionsRepository interface {
	CreateMission(ctx context.Context, mission *entities.MissionsEntity) (*entities.MissionsEntity, error)
	UpdateMission(ctx context.Context, mission *entities.MissionsEntity) (*entities.MissionsEntity, error)
	DeleteMission(ctx context.Context, ID int) (*entities.MissionsEntity, error)

	ReadMission(ctx context.Context, ID int) (*entities.MissionsEntity, error)
}
