package repositories

import "github.com/jansen44/nagasu-kujira/core/entities"

type IProjectsRepository interface {
	CreateProject(project *entities.ProjectsEntity) (*entities.ProjectsEntity, error)
	ReadProjects() ([]entities.ProjectsEntity, error)
	UpdateProject(project *entities.ProjectsEntity) (*entities.ProjectsEntity, error)
	DeleteProject(ID int) (*entities.ProjectsEntity, error)

	ReadProject(projectID int) (*entities.ProjectsEntity, error)
}
