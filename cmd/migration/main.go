package main

import (
	"flag"
	"github.com/OvictorVieira/transact.ease/internal/config"
	"github.com/OvictorVieira/transact.ease/internal/constants"
	"github.com/OvictorVieira/transact.ease/pkg/helpers"
	"github.com/OvictorVieira/transact.ease/pkg/logger"
	"github.com/sirupsen/logrus"
)

var (
	up   bool
	down bool
)

func init() {
	if err := config.InitializeAppConfig(); err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
	}
	logger.Info("configuration loaded", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
}

func main() {
	flag.BoolVar(&up, "up", false, "involves creating new tables, columns, or other database structures")
	flag.BoolVar(&down, "down", false, "involves dropping tables, columns, or other structures")
	flag.Parse()

	db, err := config.SetupPostgresConnection()
	if err != nil {
		logger.Panic(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryMigration})
	}
	defer db.Close()

	if up {
		err = helpers.Migrate(db, "up")
		if err != nil {
			logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryMigration})
		}
	}

	if down {
		err = helpers.Migrate(db, "down")
		if err != nil {
			logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryMigration})
		}
	}
}
