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
	// AutomaticEnv tells Viper to check environment variables
	// for all the keys as it tries to find values for your config struct.
	viper.AutomaticEnv()

	// Optionally you can also set a prefix for environment variables
	// viper.SetEnvPrefix("APP")

	// If running locally, attempt to load the configuration from a file.
	if _, local := os.LookupEnv("LOCAL"); local {
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.AddConfigPath("../config")
		viper.AddConfigPath("/")
		viper.AllowEmptyEnv(true)

		if err := viper.ReadInConfig(); err != nil {
			LOGGER.Error("error when try to load configs: "+err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})
			return constants.ErrLoadConfig
		}
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return constants.ErrParseConfig
	}

	LOGGER.Info("configs env: "+AppConfig.Environment, logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})

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
