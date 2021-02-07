package models

import "github.com/jansen44/nagasu-kujira/core/entities"

type SliceProjectResponse struct {
	Code  int                       `json:"code"`
	Data  []entities.ProjectsEntity `json:"data"`
	Error error                     `json:"error"`
}

type ProjectResponse struct {
	Code  int                      `json:"code"`
	Data  *entities.ProjectsEntity `json:"data"`
	Error error                    `json:"error"`
}

type NewProjectSerializer struct {
	Name string `json:"name"`
}

type UpdateProjectSerializer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
