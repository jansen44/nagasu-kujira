package util

import "github.com/jansen44/nagasu-kujira/core/repositories"

type DataSources string
type Drivers string

const (
	MySqlDataSource DataSources = "mysql"
)

const (
	RestAPIDriver Drivers = "rest"
)

type Config struct {
	DataSource DataSources
	Driver     Drivers

	// Repositories ==================
	ProjectRepository repositories.IProjectsRepository
}
