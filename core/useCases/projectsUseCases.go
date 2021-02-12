package useCases

import (
	"context"
	"github.com/jansen44/nagasu-kujira/core/entities"
	"github.com/jansen44/nagasu-kujira/core/repositories"
)

type IProjectsUseCases interface {
	ListProjects(ctx context.Context) ([]entities.ProjectsEntity, error)
	GetProjectInfo(ctx context.Context, ID int) (*entities.ProjectsEntity, error)
	AddNewProject(ctx context.Context, name string) (*entities.ProjectsEntity, error)
	UpdateOneProject(ctx context.Context, name string, ID int) (*entities.ProjectsEntity, error)
	RemoveOneProject(ctx context.Context, ID int) (*entities.ProjectsEntity, error)
}

type ProjectsUseCases struct {
	projectsRepository repositories.IProjectsRepository
}

func NewProjectsUseCases(projectsRepository repositories.IProjectsRepository) IProjectsUseCases {
	return &ProjectsUseCases{
		projectsRepository: projectsRepository,
	}
}

func (p ProjectsUseCases) ListProjects(ctx context.Context) ([]entities.ProjectsEntity, error) {
	return p.projectsRepository.ReadProjects(ctx)
}

func (p ProjectsUseCases) AddNewProject(ctx context.Context, name string) (*entities.ProjectsEntity, error) {
	return p.projectsRepository.CreateProject(ctx, entities.NewProjectsEntity(name, 0))
}

func (p ProjectsUseCases) UpdateOneProject(ctx context.Context, name string, ID int) (*entities.ProjectsEntity, error) {
	return p.projectsRepository.UpdateProject(ctx, entities.NewProjectsEntity(name, ID))
}

func (p ProjectsUseCases) RemoveOneProject(ctx context.Context, ID int) (*entities.ProjectsEntity, error) {
	return p.projectsRepository.DeleteProject(ctx, ID)
}

func (p ProjectsUseCases) GetProjectInfo(ctx context.Context, ID int) (*entities.ProjectsEntity, error) {
	return p.projectsRepository.ReadProject(ctx, ID)
}
