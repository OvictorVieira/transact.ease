package config

import (
	"github.com/OvictorVieira/transact.ease/internal/constants"
	LOGGER "github.com/OvictorVieira/transact.ease/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	Port        int    `mapstructure:"PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`
	Debug       bool   `mapstructure:"DEBUG"`

	DBPostgreDriver  string `mapstructure:"DB_POSTGRE_DRIVER"`
	DBPostgreDsn     string `mapstructure:"DB_POSTGRE_DSN"`
	DBPostgreURL     string `mapstructure:"DB_POSTGRE_URL"`
	DBPostgreDsnTest string `mapstructure:"DB_POSTGRE_DSN_TEST"`
}

func InitializeAppConfig() error {
	viper.SetConfigName(".env") // allow directly reading from .env file
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("internal/config")
	viper.AddConfigPath("/")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		LOGGER.Error("error when try to load configs: "+err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})
		return constants.ErrLoadConfig
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		return constants.ErrParseConfig
	}

	// check
	if AppConfig.Port == 0 || AppConfig.Environment == "" || AppConfig.DBPostgreDriver == "" {
		return constants.ErrEmptyVar
	}

	switch AppConfig.Environment {
	case constants.EnvironmentDevelopment:
		if AppConfig.DBPostgreDsn == "" {
			return constants.ErrEmptyVar
		}
	case constants.EnvironmentProduction:
		if AppConfig.DBPostgreURL == "" {
			return constants.ErrEmptyVar
		}
	}

	return nil
}
