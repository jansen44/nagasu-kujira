package models

import "github.com/jansen44/nagasu-kujira/core/entities"

type SliceMissionResponse struct {
	Code  int                       `json:"code"`
	Data  []entities.MissionsEntity `json:"data"`
	Error error                     `json:"error"`
}

type MissionResponse struct {
	Code  int                      `json:"code"`
	Data  *entities.MissionsEntity `json:"data"`
	Error error                    `json:"error"`
}

type NewMissionSerializer struct {
	ProjectID int    `json:"projectId"`
	Name      string `json:"name"`
}

type UpdateMissionSerializer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
