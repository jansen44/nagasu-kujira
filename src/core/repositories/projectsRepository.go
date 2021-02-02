package repositories

import "github.com/jansen44/nagasu-kujira/src/core/entities"

type IProjectsRepository interface {
	// CRUD ==========================================================
	CreateProject(project *entities.ProjectsEntity) (*entities.ProjectsEntity, error)
	ReadProjects() ([]entities.ProjectsEntity, error)
	UpdateProject(project *entities.ProjectsEntity) (*entities.ProjectsEntity, error)
	DeleteProject(ID int32) (*entities.ProjectsEntity, error)
}
