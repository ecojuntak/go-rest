package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ecojuntak/go-rest/config"
	"github.com/ecojuntak/go-rest/controller"
	"github.com/ecojuntak/go-rest/model"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.LoadConfig()
	model.InitDatabase()
}

func TestGetAllUsers(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAllUsers)
	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
