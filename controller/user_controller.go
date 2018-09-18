package controller

import (
	"net/http"

	"github.com/ecojuntak/go-rest/model"
)

var user model.User

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := user.GetAll()
	RespondWithJSON(w, http.StatusOK, users)
}
