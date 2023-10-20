package helpers

import (
	"fmt"
	"github.com/OvictorVieira/transact.ease/cmd/helpers"
	"github.com/OvictorVieira/transact.ease/internal/config"
	"github.com/OvictorVieira/transact.ease/internal/constants"
	"github.com/OvictorVieira/transact.ease/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sort"
)

const (
	dirTest = "../../cmd/migration/migrations"
	dir     = "cmd/migration/migrations"
)

func Migrate(db *sqlx.DB, action string) (err error) {
	logger.InfoF("running migration [%s]", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryMigration}, action)

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	migrationsDirectory := dir

	if config.AppConfig.Environment == constants.EnvironmentTest {
		migrationsDirectory = dirTest
	}

	files, err := filepath.Glob(filepath.Join(cwd, migrationsDirectory, fmt.Sprintf("*.%s.sql", action)))
	if err != nil {
		return constants.ErrWhenGetFiles
	}

	sort.Slice(files, func(i, j int) bool {
		return helpers.ExtractLeadingNumber(files[i]) < helpers.ExtractLeadingNumber(files[j])
	})

	for _, file := range files {
		logger.Info("Executing migration", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryMigration, constants.LoggerFile: file})
		data, err := os.ReadFile(file)
		if err != nil {
			return constants.ErrWhenReadFiles
		}

		_, err = db.Exec(string(data))
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("error when exec query in file: %v", file)
		}
	}

	logger.InfoF("migration [%s] success", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryMigration}, action)

	return
}
