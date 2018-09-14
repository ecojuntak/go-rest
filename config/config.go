package config

import (
	"github.com/ecojuntak/go-rest/util"
)

func GetConfigFilePath() string {
	rootPath := util.GetRootFolderPath()
	const configDirectory = "config/"
	const configFileName = "config.yaml"

	return rootPath + configDirectory + configFileName
}
