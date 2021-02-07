package main

import (
	"github.com/jansen44/nagasu-kujira/drivers/rest"
	"github.com/jansen44/nagasu-kujira/infra/mysql"
	"github.com/jansen44/nagasu-kujira/util"
	"github.com/sirupsen/logrus"
)

func initRepositories(config *util.Config) error {
	logrus.Info("## Initializing Repositories...")
	switch config.DataSource {
	case util.MySqlDataSource:
		err := mysql.InitRepositories(config)
		if err != nil {
			return err
		}
	}
	return nil
}

func initDrivers(config *util.Config) error {
	logrus.Info("## Initializing Drivers...")
	switch config.Driver {
	case util.RestAPIDriver:
		client := rest.NewRestAPIClient(config)
		client.Start()
	}
	return nil
}

func main() {
	config := util.Config{DataSource: util.MySqlDataSource, Driver: util.RestAPIDriver}
	err := initRepositories(&config)
	if err != nil {
		panic(err)
	}
	err = initDrivers(&config)
	if err != nil {
		panic(err)
	}
}
