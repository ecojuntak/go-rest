package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ecojuntak/go-rest/config"
	"github.com/ecojuntak/go-rest/model"
	"github.com/ecojuntak/go-rest/router"
	"github.com/urfave/cli"
)

func InitConfig() (err error) {
	err = config.LoadConfig()
	return
}

func startServer() {
	err := InitConfig()
	if err != nil {
		log.Fatal("Error while loading the configuration")
		log.Fatal(err.Error())
	}

	model.InitDatabase()

	server := &http.Server{
		Handler: router.InitRouter(),
		Addr:    config.GetAddress(),
	}

	log.Println("Server running on", config.GetAddress())
	log.Println(server.ListenAndServe())
}

func main() {
	cliApp := cli.NewApp()
	cliApp.Name = "GO-REST"
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		{
			Name:        "migrate",
			Description: "Run database migration",
			Action: func(c *cli.Context) error {
				err := model.Migrate()
				return err
			},
		},
		{
			Name:        "start",
			Description: "Start REST API Server",
			Action: func(c *cli.Context) error {
				startServer()
				return nil
			},
		},
	}

	cliApp.Run(os.Args)
}
