package main

import (
	"fmt"
	"github.com/OvictorVieira/transact.ease/cmd/helpers"
	"github.com/OvictorVieira/transact.ease/internal/config"
	"github.com/OvictorVieira/transact.ease/internal/constants"
	"github.com/OvictorVieira/transact.ease/pkg/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sort"
)

const (
	dir = "cmd/seed/seeds"
)

func init() {
	if err := config.InitializeAppConfig(); err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
	}
	logger.Info("configuration loaded", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
}

func main() {
	db, err := config.SetupPostgresConnection()
	if err != nil {
		logger.Panic(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})
	}
	defer db.Close()

	err = seed(db)

	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})
	}
}

func seed(db *sqlx.DB) (err error) {
	logger.Info("Seeding...", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	files, err := filepath.Glob(filepath.Join(cwd, dir, fmt.Sprintf("*.sql")))
	if err != nil {
		return constants.ErrWhenGetFiles
	}

	sort.Slice(files, func(i, j int) bool {
		return helpers.ExtractLeadingNumber(files[i]) < helpers.ExtractLeadingNumber(files[j])
	})

	for _, file := range files {
		logger.Info("Executing seed", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder, constants.LoggerFile: file})
		data, err := os.ReadFile(file)
		if err != nil {
			return constants.ErrWhenReadFiles
		}

		_, err = db.Exec(string(data))
		if err != nil {
			fmt.Println(err.Error())
			return fmt.Errorf("error when exec query in file: %v. Error: %s", file, err.Error())
		}
	}

	logger.InfoF("seeds executed with success", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})

	return
}
