package config

import (
	"github.com/OvictorVieira/transact.ease/internal/constants"
	"github.com/jmoiron/sqlx"
	"time"
)

func SetupPostgresConnection() (*sqlx.DB, error) {
	var dsn string
	switch AppConfig.Environment {
	case constants.EnvironmentTest:
		dsn = AppConfig.DBPostgreDsnTest
	case constants.EnvironmentDevelopment:
		dsn = AppConfig.DBPostgreDsn
	case constants.EnvironmentProduction:
		dsn = AppConfig.DBPostgreURL
	}

	sqlxConfig := SQLXConfig{
		DriverName:     AppConfig.DBPostgreDriver,
		DataSourceName: dsn,
		MaxOpenConns:   100,
		MaxIdleConns:   10,
		MaxLifetime:    15 * time.Minute,
	}

	conn, err := sqlxConfig.InitializeSQLXDatabase()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
