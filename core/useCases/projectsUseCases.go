package useCases

import (
	"github.com/jansen44/nagasu-kujira/core/entities"
	"github.com/jansen44/nagasu-kujira/core/repositories"
)

type IProjectsUseCases interface {
	ListProjects() ([]entities.ProjectsEntity, error)
	AddNewProject(name string) (*entities.ProjectsEntity, error)
	UpdateOneProject(name string, ID int) (*entities.ProjectsEntity, error)
	RemoveOneProject(ID int) (*entities.ProjectsEntity, error)
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
	return p.projectsRepository.ReadProjects()
}

func (p ProjectsUseCases) AddNewProject(name string) (*entities.ProjectsEntity, error) {
	return p.projectsRepository.CreateProject(entities.NewProjectsEntity(name, 0))
}

func (p ProjectsUseCases) UpdateOneProject(name string, ID int) (*entities.ProjectsEntity, error) {
	return p.projectsRepository.UpdateProject(entities.NewProjectsEntity(name, ID))
}

func (p ProjectsUseCases) RemoveOneProject(ID int) (*entities.ProjectsEntity, error) {
	return p.projectsRepository.DeleteProject(ID)
}
