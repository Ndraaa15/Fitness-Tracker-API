package config

import (
	"github/Ndraaa15/fitness-tracker-api/pkg/utils/errors"

	"github.com/spf13/viper"
)

const (
	configName   = "config"
	configType   = "json"
	configPath   = "."
	configPrefix = "config"
)

func InitializeViper() (*viper.Viper, error) {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	viper.SetEnvPrefix(configPrefix)

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.ErrConfigFileNotFound
		} else {
			return nil, err
		}
	}

	return viper.GetViper(), nil
}
