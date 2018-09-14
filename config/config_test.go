package config_test

import (
	"os"
	"strings"
	"testing"

	"github.com/ecojuntak/go-rest/config"
	"github.com/stretchr/testify/assert"
)

func TestGetRootPath(t *testing.T) {
	rootPath := config.GetConfigFilePath()
	assert.Equal(t, strings.Contains(rootPath, "config/config.yaml"), true)
}

func TestConfigFileIsExist(t *testing.T) {
	fullPathConfigFile := config.GetConfigFilePath()
	_, err := os.Stat(fullPathConfigFile)
	assert.Equal(t, nil, err)
}
