package mysql

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jansen44/nagasu-kujira/util"
	"github.com/sirupsen/logrus"
)

func InitRepositories(config *util.Config) error {
	logrus.Info("## Initializing MySQL Repositories!")
	mysqlContext := context.Background()

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/nagasu?parseTime=true")
	if err != nil {
		return err
	}

	config.ProjectRepository = NewProjectsMySQL(db, mysqlContext)
	return nil
}
