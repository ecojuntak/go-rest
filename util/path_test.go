package util_test

import (
	"os"
	"testing"

	"github.com/ecojuntak/go-rest/util"
	"github.com/stretchr/testify/assert"
)

func TestGetRootPath_ExpectedSuccess(t *testing.T) {
	rootProjectPath := util.GetRootFolderPath()
	_, err := os.Stat(rootProjectPath)
	assert.Equal(t, nil, err)
}
