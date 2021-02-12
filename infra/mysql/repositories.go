package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jansen44/nagasu-kujira/util"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)

func InitRepositories(config *util.Config) error {
	logrus.Info("## Initializing MySQL DB Connection...")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/nagasu?parseTime=true")
	if err != nil {
		return err
	}

	if config.ShouldMigrate {
		logrus.Info("## Handling MySQL migrations...")
		migrations := &migrate.FileMigrationSource{
			Dir: "infra/mysql/migrations",
		}
		_, err = migrate.Exec(db, "mysql", migrations, migrate.Up)
		if err != nil {
			return err
		}
	}

	logrus.Info("## Initializing MySQL Repositories...")
	config.ProjectRepository = NewProjectsMySQL(db)
	config.MissionRepository = NewMissionsMySQL(db)
	config.TaskRepository = NewTasksMySQL(db)

	return nil
}
