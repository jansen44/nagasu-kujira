package models

import "github.com/jansen44/nagasu-kujira/core/entities"

type SliceTaskResponse struct {
	Code  int                    `json:"code"`
	Data  []entities.TasksEntity `json:"data"`
	Error error                  `json:"error"`
}

type TaskResponse struct {
	Code  int                   `json:"code"`
	Data  *entities.TasksEntity `json:"data"`
	Error error                 `json:"error"`
}

type NewTaskSerializer struct {
	MissionID   int    `json:"missionId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StatusID    int64  `json:"statusId"`
}

type UpdateTaskSerializer struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StatusID    int64  `json:"statusId"`
}
