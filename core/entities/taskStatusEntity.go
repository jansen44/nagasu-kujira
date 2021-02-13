package entities

type TaskStatusEntity struct {
	ID        int64  `json:"id"`
	MissionID int    `json:"missionId"`
	Name      string `json:"name"`
	Order     int    `json:"order"`
}

type DefaultTaskStatusNames string

const (
	ToDo  DefaultTaskStatusNames = "To Do"
	Doing                        = "Doing"
	Done                         = "Done"
)

func NewTaskStatusEntity(Name string, MissionID int, Order int, ID int64) *TaskStatusEntity {
	return &TaskStatusEntity{
		ID,
		MissionID,
		Name,
		Order,
	}
}

func GetDefaultTaskStatuses() []DefaultTaskStatusNames {
	initialStatus := []DefaultTaskStatusNames{ToDo, Doing, Done}
	return initialStatus
}
