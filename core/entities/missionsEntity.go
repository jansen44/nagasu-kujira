package entities

type MissionsEntity struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"projectId"`
	Name      string `json:"name"`
}

func NewMissionsEntity(Name string, ProjectID, ID int) *MissionsEntity {
	return &MissionsEntity{
		ID:        ID,
		ProjectID: ProjectID,
		Name:      Name,
	}
}
