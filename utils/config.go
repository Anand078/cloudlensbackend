package utils

import (
	"backend/logging"

	"github.com/spf13/viper"
)

func GetEnvironmentVars(propName string) string {
	viper.SetConfigFile("app_dev.env")
	err := viper.ReadInConfig()

	if err != nil {
		logging.Logger.Errorf(err.Error())
	}

	value, ok := viper.Get(propName).(string)
	if !ok {
		logging.Logger.Error(ok)
	}

	return value
}
