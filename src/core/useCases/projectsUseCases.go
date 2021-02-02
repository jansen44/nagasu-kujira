package useCases

import (
	"github.com/jansen44/nagasu-kujira/src/core/entities"
	"github.com/jansen44/nagasu-kujira/src/core/repositories"
)

type IProjectsUseCases interface {
	ListProjects() ([]entities.ProjectsEntity, error)
	AddNewProject(name string) (*entities.ProjectsEntity, error)
	UpdateOneProject(name string, ID int32) (*entities.ProjectsEntity, error)
	RemoveOneProject(ID int32) (*entities.ProjectsEntity, error)
}

type ProjectsUseCases struct {
	projectsRepository repositories.IProjectsRepository
}

func NewProjectsUseCases(projectsRepository repositories.IProjectsRepository) IProjectsUseCases {
	return &ProjectsUseCases{
		projectsRepository: projectsRepository,
	}
}

func (p ProjectsUseCases) ListProjects() ([]entities.ProjectsEntity, error) {
	panic("implement me")
}

func (p ProjectsUseCases) AddNewProject(name string) (*entities.ProjectsEntity, error) {
	panic("implement me")
}

func (p ProjectsUseCases) UpdateOneProject(name string, ID int32) (*entities.ProjectsEntity, error) {
	panic("implement me")
}

func (p ProjectsUseCases) RemoveOneProject(ID int32) (*entities.ProjectsEntity, error) {
	panic("implement me")
}
