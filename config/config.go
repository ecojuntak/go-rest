package config

import (
	"github.com/ecojuntak/go-rest/util"
	"github.com/spf13/viper"
)

func GetConfigFilePath() string {
	rootPath := util.GetRootFolderPath()
	const configDirectory = "config/"
	const configFileName = "config.yaml"

	return rootPath + configDirectory + configFileName
}

func LoadConfig() (err error) {
	configFilePath := GetConfigFilePath()
	viper.SetConfigFile(configFilePath)
	err = viper.ReadInConfig()

	return
}
