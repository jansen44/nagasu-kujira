package mysql

import (
	"context"
	"database/sql"
	"github.com/jansen44/nagasu-kujira/core/entities"
	"github.com/jansen44/nagasu-kujira/core/repositories"
	"github.com/jansen44/nagasu-kujira/infra/mysql/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ProjectsMySQL struct {
	db  *sql.DB
	ctx context.Context
}

// Util ===============================================================================
func NewProjectsMySQL(db *sql.DB, ctx context.Context) repositories.IProjectsRepository {
	return &ProjectsMySQL{db, ctx}
}

func projectModelToEntity(model *models.Project) *entities.ProjectsEntity {
	project := entities.NewProjectsEntity(model.Name, model.ID)
	if model.R != nil {
		if model.R.Missions != nil {
			project.Missions = missionModelSliceToPointerEntity(model.R.Missions)
		}
	}
	return project
}

func projectModelSliceToEntity(projectSlice models.ProjectSlice) []entities.ProjectsEntity {
	projects := make([]entities.ProjectsEntity, 0)
	for _, project := range projectSlice {
		projects = append(projects, *projectModelToEntity(project))
	}
	return projects
}

func projectEntityToModel(entity *entities.ProjectsEntity, model *models.Project) {
	model.ID = entity.ID
	model.Name = entity.Name
}

// Repository Methods ==================================================================
func (p ProjectsMySQL) CreateProject(project *entities.ProjectsEntity) (*entities.ProjectsEntity, error) {
	model := models.Project{Name: project.Name}
	err := model.Insert(p.ctx, p.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return projectModelToEntity(&model), nil
}

func (p ProjectsMySQL) ReadProjects() ([]entities.ProjectsEntity, error) {
	modelProjects, err := models.Projects().All(p.ctx, p.db)
	if err != nil {
		return nil, err
	}
	return projectModelSliceToEntity(modelProjects), nil
}

func (p ProjectsMySQL) UpdateProject(project *entities.ProjectsEntity) (*entities.ProjectsEntity, error) {
	model, err := models.FindProject(p.ctx, p.db, project.ID)
	if err != nil {
		return nil, err
	}
	projectEntityToModel(project, model)
	_, err = model.Update(p.ctx, p.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return projectModelToEntity(model), nil
}

func (p ProjectsMySQL) DeleteProject(ID int) (*entities.ProjectsEntity, error) {
	model, err := models.FindProject(p.ctx, p.db, ID)
	if err != nil {
		return nil, err
	}
	_, err = model.Delete(p.ctx, p.db)
	if err != nil {
		return nil, err
	}
	return projectModelToEntity(model), nil
}

func (p ProjectsMySQL) ReadProject(projectID int) (*entities.ProjectsEntity, error) {
	project, err := models.Projects(qm.Load(models.ProjectRels.Missions), qm.Where("id=?", projectID)).One(p.ctx, p.db)
	if err != nil {
		return nil, err
	}
	return projectModelToEntity(project), nil
}
