package entities

type ProjectsEntity struct {
	ID   *int32 `json:"id"`
	Name string `json:"name"`
}

func NewProjectsEntity(name string, ID *int32) *ProjectsEntity {
	return &ProjectsEntity{
		ID:   ID,
		Name: name,
	}
}
