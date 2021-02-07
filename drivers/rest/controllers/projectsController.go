package controllers

import (
	"encoding/json"
	"github.com/jansen44/nagasu-kujira/drivers/rest/models"
	"net/http"
	"strconv"
)

func (api RestAPIController) GetProjects(w http.ResponseWriter, _ *http.Request) {
	projects, err := api.projectsUseCases.ListProjects()
	if err != nil {
		_ = json.NewEncoder(w).Encode(models.SliceProjectResponse{Code: 500, Error: err})
	} else {
		_ = json.NewEncoder(w).Encode(models.SliceProjectResponse{Code: 200, Data: projects})
	}
}

func (api RestAPIController) PostProject(w http.ResponseWriter, r *http.Request) {
	var projectSerializer models.NewProjectSerializer
	_ = json.NewDecoder(r.Body).Decode(&projectSerializer)
	project, err := api.projectsUseCases.AddNewProject(projectSerializer.Name)
	if err != nil {
		_ = json.NewEncoder(w).Encode(models.ProjectResponse{Code: 500, Error: err})
	} else {
		_ = json.NewEncoder(w).Encode(models.ProjectResponse{Code: 200, Data: project})
	}
}

func (api RestAPIController) PutProject(w http.ResponseWriter, r *http.Request) {
	var projectSerializer models.UpdateProjectSerializer
	_ = json.NewDecoder(r.Body).Decode(&projectSerializer)
	project, err := api.projectsUseCases.UpdateOneProject(projectSerializer.Name, projectSerializer.ID)
	if err != nil {
		_ = json.NewEncoder(w).Encode(models.ProjectResponse{Code: 500, Error: err})
	} else {
		_ = json.NewEncoder(w).Encode(models.ProjectResponse{Code: 200, Data: project})
	}
}

func (api RestAPIController) DeleteProject(w http.ResponseWriter, r *http.Request) {
	idKeys, ok := r.URL.Query()["id"]
	if !ok || len(idKeys) == 0 {
		_ = json.NewEncoder(w).Encode(models.ProjectResponse{Code: 403})
		return
	}

	id, err := strconv.Atoi(idKeys[0])
	if err != nil {
		_ = json.NewEncoder(w).Encode(models.ProjectResponse{Code: 403})
		return
	}

	project, err := api.projectsUseCases.RemoveOneProject(id)
	if err != nil {
		_ = json.NewEncoder(w).Encode(models.ProjectResponse{Code: 500, Error: err})
	} else {
		_ = json.NewEncoder(w).Encode(models.ProjectResponse{Code: 200, Data: project})
	}
}
