package controllers

import (
	"context"
	"encoding/json"
	"github.com/jansen44/nagasu-kujira/drivers/rest/models"
	"net/http"
	"strconv"
)

func (api RestAPIController) PortTask(w http.ResponseWriter, r *http.Request) {
	var taskSerializer models.NewTaskSerializer
	_ = json.NewDecoder(r.Body).Decode(&taskSerializer)
	task, err := api.tasksUseCases.AddNewTask(context.Background(), taskSerializer.Name, taskSerializer.Description, taskSerializer.MissionID)
	if err != nil {
		_ = json.NewEncoder(w).Encode(models.TaskResponse{Code: 500, Error: err})
	} else {
		_ = json.NewEncoder(w).Encode(models.TaskResponse{Code: 200, Data: task})
	}
}

func (api RestAPIController) PutTask(w http.ResponseWriter, r *http.Request) {
	var taskSerializer models.UpdateTaskSerializer
	_ = json.NewDecoder(r.Body).Decode(&taskSerializer)
	task, err := api.tasksUseCases.UpdateOneTask(context.Background(), taskSerializer.Name, taskSerializer.Description, taskSerializer.ID)
	if err != nil {
		_ = json.NewEncoder(w).Encode(models.TaskResponse{Code: 500, Error: err})
	} else {
		_ = json.NewEncoder(w).Encode(models.TaskResponse{Code: 200, Data: task})
	}
}

func (api RestAPIController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idKeys, ok := r.URL.Query()["id"]
	if !ok || len(idKeys) == 0 {
		_ = json.NewEncoder(w).Encode(models.TaskResponse{Code: 403})
		return
	}
	id, err := strconv.ParseInt(idKeys[0], 10, 64)
	if err != nil {
		_ = json.NewEncoder(w).Encode(models.TaskResponse{Code: 403})
		return
	}
	task, err := api.tasksUseCases.RemoveOneTask(context.Background(), id)
	if err != nil {
		_ = json.NewEncoder(w).Encode(models.TaskResponse{Code: 500, Error: err})
	} else {
		_ = json.NewEncoder(w).Encode(models.TaskResponse{Code: 200, Data: task})
	}
}
