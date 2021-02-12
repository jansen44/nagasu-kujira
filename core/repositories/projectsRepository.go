package repositories

import (
	"context"
	"github.com/jansen44/nagasu-kujira/core/entities"
)

type IProjectsRepository interface {
	CreateProject(ctx context.Context, project *entities.ProjectsEntity) (*entities.ProjectsEntity, error)
	ReadProjects(ctx context.Context) ([]entities.ProjectsEntity, error)
	UpdateProject(ctx context.Context, project *entities.ProjectsEntity) (*entities.ProjectsEntity, error)
	DeleteProject(ctx context.Context, ID int) (*entities.ProjectsEntity, error)

	ReadProject(ctx context.Context, projectID int) (*entities.ProjectsEntity, error)
}
