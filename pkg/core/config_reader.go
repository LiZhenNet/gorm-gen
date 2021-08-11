package core

import (
	"errors"

	"github.com/spf13/viper"
)

func ReadConfig(config string) (GeneratorConfig, error) {
	viper.SetConfigFile(config)
	generatorConfig := GeneratorConfig{}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return generatorConfig, errors.New("config file not found")
		} else {
			return generatorConfig, errors.New("read config error,msg:" + err.Error())
		}
	}
	err := viper.Unmarshal(&generatorConfig)
	if err != nil {
		return generatorConfig, errors.New("parse config error,msg:" + err.Error())
	}
	return generatorConfig, nil
}
