package util

import (
	"flag"
	"github.com/jansen44/nagasu-kujira/core/repositories"
	"os"
)

type Config struct {
	fs *flag.FlagSet

	Storage  string `yaml:"data-source"`
	Driver   string `yaml:"driver"`
	RestPort string `yaml:"port"`

	// Repositories ==================
	ProjectRepository repositories.IProjectsRepository
	MissionRepository repositories.IMissionsRepository
	TaskRepository    repositories.ITasksRepository
}

func NewConfig() (*Config, error) {
	config := Config{fs: flag.NewFlagSet("kujira", flag.ExitOnError)}
	config.fs.StringVar(&config.Storage, "storage", "mysql", "set what storage will be used (defaults to 'mysql')")
	config.fs.StringVar(&config.Driver, "driver", "rest", "set what driver will be initiated (defaults to 'rest')")
	config.fs.StringVar(&config.RestPort, "restPort", "8080", "set what port restApi should listen on (defaults to 8080")
	err := config.fs.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}
	return &config, nil
}
