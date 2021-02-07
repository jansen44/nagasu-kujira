package entities

type ProjectsEntity struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Missions []MissionsEntity `json:"missions"`
}

func NewProjectsEntity(Name string, ID int) *ProjectsEntity {
	return &ProjectsEntity{
		ID:       ID,
		Name:     Name,
		Missions: make([]MissionsEntity, 0),
	}
}
