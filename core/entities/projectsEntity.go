package entities

type ProjectsEntity struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Missions []MissionsEntity `json:"missions"`
}
