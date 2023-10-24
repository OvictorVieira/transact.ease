package config

import (
	"github.com/OvictorVieira/transact.ease/internal/constants"
	LOGGER "github.com/OvictorVieira/transact.ease/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var AppConfig Config

type Config struct {
	Port        int    `mapstructure:"PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`
	Debug       bool   `mapstructure:"DEBUG"`

	DBPostgreDriver string `mapstructure:"DB_POSTGRE_DRIVER"`
	DBPostgreDsn    string `mapstructure:"DB_POSTGRE_DSN"`
	DBPostgreURL    string `mapstructure:"DATABASE_URL"`
}

func InitializeAppConfig() error {
	viper.AutomaticEnv()

	if _, local := os.LookupEnv("LOCAL"); local {
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.AddConfigPath("../config")
		viper.AddConfigPath("internal/config")
		viper.AddConfigPath("/")
		viper.AllowEmptyEnv(true)
	} else {
		LOGGER.InfoF("binding environment variables", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})
		viper.BindEnv("PORT")
		viper.BindEnv("ENVIRONMENT")
		viper.BindEnv("DATABASE_URL")
	}

	if err := viper.ReadInConfig(); err != nil {
		LOGGER.Error("error when try to load configs: "+err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})
		return constants.ErrLoadConfig
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		LOGGER.Error("error when try to unmarshall configs: "+err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})
		return constants.ErrParseConfig
	}

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
