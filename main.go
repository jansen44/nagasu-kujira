package main

import (
	"github.com/friendsofgo/errors"
	"github.com/jansen44/nagasu-kujira/drivers/rest"
	"github.com/jansen44/nagasu-kujira/infra/mysql"
	"github.com/jansen44/nagasu-kujira/util"
	"github.com/sirupsen/logrus"
)

func initStorage(config *util.Config) error {
	logrus.Info("## Initializing Storage...")
	switch config.Storage {
	case "mysql":
		err := mysql.InitRepositories(config)
		if err != nil {
			return err
		}
	default:
		return errors.New("storage error: no valid storage was supplied")
	}
	return nil
}

func initDriver(config *util.Config) error {
	logrus.Info("## Initializing Drivers...")
	switch config.Driver {
	case "rest":
		client := rest.NewRestAPIClient(config)
		client.Start()
	default:
		return errors.New("driver error: no valid driver was supplied")
	}
	return nil
}

func main() {
	config, err := util.NewConfig()
	if err != nil {
		panic(err)
	}
	err = initStorage(config)
	if err != nil {
		panic(err)
	}
	err = initDriver(config)
	if err != nil {
		panic(err)
	}
}
