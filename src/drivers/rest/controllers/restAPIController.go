package controllers

import (
	"github.com/jansen44/nagasu-kujira/src/core/repositories"
	"github.com/jansen44/nagasu-kujira/src/core/useCases"
	"net/http"
)

type IRestAPIController interface {
	GetProjects(w http.ResponseWriter, r *http.Request)
	PostProject(w http.ResponseWriter, r *http.Request)
	PutProject(w http.ResponseWriter, r *http.Request)
	DeleteProject(w http.ResponseWriter, r *http.Request)
}

type RestAPIController struct {
	projectsUseCases useCases.IProjectsUseCases
}

func NewRestAPIController(projectsRepository repositories.IProjectsRepository) IRestAPIController {
	return &RestAPIController{
		projectsUseCases: useCases.NewProjectsUseCases(projectsRepository),
	}
}
