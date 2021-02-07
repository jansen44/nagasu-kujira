package controllers

import (
	"github.com/jansen44/nagasu-kujira/core/repositories"
	"github.com/jansen44/nagasu-kujira/core/useCases"
	"net/http"
)

type IRestAPIController interface {
	GetProjects(w http.ResponseWriter, r *http.Request)
	PostProject(w http.ResponseWriter, r *http.Request)
	PutProject(w http.ResponseWriter, r *http.Request)
	DeleteProject(w http.ResponseWriter, r *http.Request)
	GetProject(w http.ResponseWriter, r *http.Request)

	PostMission(w http.ResponseWriter, r *http.Request)
	PutMission(w http.ResponseWriter, r *http.Request)
	DeleteMission(w http.ResponseWriter, r *http.Request)
}

type RestAPIController struct {
	projectsUseCases useCases.IProjectsUseCases
	missionsUseCases useCases.IMissionsUseCases
}

func NewRestAPIController(
	projectsRepository repositories.IProjectsRepository,
	missionsRepository repositories.IMissionsRepository,
) IRestAPIController {
	return &RestAPIController{
		projectsUseCases: useCases.NewProjectsUseCases(projectsRepository),
		missionsUseCases: useCases.NewMissionsUseCases(missionsRepository),
	}
}
