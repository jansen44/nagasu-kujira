package entities

type TasksEntity struct {
	ID          int64  `json:"id"`
	MissionID   int    `json:"missionId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewTasksEntity(Name, Description string, MissionID int, ID int64) *TasksEntity {
	return &TasksEntity{
		ID,
		MissionID,
		Name,
		Description,
	}
}
