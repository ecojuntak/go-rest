package app_test

import (
	"testing"

	"github.com/ecojuntak/go-rest/app"
	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	err := app.InitConfig()
	assert.Equal(t, nil, err)
}
