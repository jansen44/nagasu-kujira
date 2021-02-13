package util

import (
	"flag"
	"github.com/jansen44/nagasu-kujira/core/repositories"
	"os"
)

type Config struct {
	fs *flag.FlagSet

	Storage       string `yaml:"data-source"`
	Driver        string `yaml:"driver"`
	RestPort      string `yaml:"port"`
	ShouldMigrate bool   `yaml:"should-migrate"`

	// Repositories ==================
	ProjectRepository    repositories.IProjectsRepository
	MissionRepository    repositories.IMissionsRepository
	TaskRepository       repositories.ITasksRepository
	TaskStatusRepository repositories.ITaskStatusRepository
}

func NewConfig() (*Config, error) {
	config := Config{fs: flag.NewFlagSet("kujira", flag.ExitOnError)}
	config.fs.StringVar(&config.Storage, "storage", "mysql", "set what storage will be used")
	config.fs.StringVar(&config.Driver, "driver", "rest", "set what driver will be initiated")
	config.fs.StringVar(&config.RestPort, "restPort", "8080", "set what port restApi should listen on")
	config.fs.BoolVar(&config.ShouldMigrate, "migrate", false, "set if the storage should execute migrations")
	err := config.fs.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}
	return &config, nil
}
