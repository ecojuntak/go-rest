package router

import (
	"github.com/ecojuntak/go-rest/controller"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.GetAllUsers).Methods("GET")

	return router
}
