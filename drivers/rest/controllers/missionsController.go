package controllers

import (
	"encoding/json"
	"github.com/jansen44/nagasu-kujira/drivers/rest/models"
	"net/http"
	"strconv"
)

func (api RestAPIController) PostMission(w http.ResponseWriter, r *http.Request) {
	var missionSerializer models.NewMissionSerializer
	_ = json.NewDecoder(r.Body).Decode(&missionSerializer)
	mission, err := api.missionsUseCases.AddNewMission(missionSerializer.Name, missionSerializer.ProjectID)
	if err != nil {
		_ = json.NewEncoder(w).Encode(models.MissionResponse{Code: 500, Error: err})
	} else {
		_ = json.NewEncoder(w).Encode(models.MissionResponse{Code: 200, Data: mission})
	}
}

func (api RestAPIController) PutMission(w http.ResponseWriter, r *http.Request) {
	var missionSerializer models.UpdateMissionSerializer
	_ = json.NewDecoder(r.Body).Decode(&missionSerializer)
	mission, err := api.missionsUseCases.UpdateOneMission(missionSerializer.Name, missionSerializer.ID)
	if err != nil {
		_ = json.NewEncoder(w).Encode(models.MissionResponse{Code: 500, Error: err})
	} else {
		_ = json.NewEncoder(w).Encode(models.MissionResponse{Code: 200, Data: mission})
	}
}

func (api RestAPIController) DeleteMission(w http.ResponseWriter, r *http.Request) {
	idKeys, ok := r.URL.Query()["id"]
	if !ok || len(idKeys) == 0 {
		_ = json.NewEncoder(w).Encode(models.MissionResponse{Code: 403})
		return
	}
	id, err := strconv.Atoi(idKeys[0])
	if err != nil {
		_ = json.NewEncoder(w).Encode(models.MissionResponse{Code: 403})
		return
	}
	mission, err := api.missionsUseCases.RemoveOneMission(id)
	if err != nil {
		_ = json.NewEncoder(w).Encode(models.MissionResponse{Code: 500, Error: err})
	} else {
		_ = json.NewEncoder(w).Encode(models.MissionResponse{Code: 200, Data: mission})
	}
}
