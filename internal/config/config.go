package config

import (
	"github.com/OvictorVieira/transact.ease/internal/constants"
	LOGGER "github.com/OvictorVieira/transact.ease/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"strings"
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

func BindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			BindEnvs(v.Interface(), append(parts, tv)...)
		default:
			viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}

func InitializeAppConfig() error {
	if _, local := os.LookupEnv("LOCAL"); local {
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.AddConfigPath("../config")
		viper.AddConfigPath("internal/config")
		viper.AddConfigPath("/")
		viper.AllowEmptyEnv(true)

		if err := viper.ReadInConfig(); err != nil {
			LOGGER.Error("error when try to load configs: "+err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})
			return constants.ErrLoadConfig
		}
	} else {
		BindEnvs(AppConfig)
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
