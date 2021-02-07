package util

import "github.com/jansen44/nagasu-kujira/core/repositories"

type DataSources string

const (
	MySqlDataSource DataSources = "mysql"
)

type Config struct {
	DataSource DataSources

	// Repositories ==================
	ProjectRepository repositories.IProjectsRepository
}
