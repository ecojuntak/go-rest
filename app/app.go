package app

import (
	"log"

	"github.com/ecojuntak/go-rest/config"
)

func InitConfig() (err error) {
	err = config.LoadConfig()
	return
}

func Run() {
	err := InitConfig()
	if err != nil {
		log.Fatal("Error while loading the configuration")
		log.Fatal(err.Error())
	}
}
