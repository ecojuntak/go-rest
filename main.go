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
	model.InitDatabase()

	server := &http.Server{
		Handler: router.InitRouter(),
		Addr:    config.GetAddress(),
	}

	log.Println("Server running on", config.GetAddress())
	log.Println(server.ListenAndServe())
}

func main() {
	err := InitConfig()
	if err != nil {
		log.Fatal("Error while loading the configuration")
		log.Fatal(err.Error())
	}

	cliApp := cli.NewApp()
	cliApp.Name = "GO-REST"
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		{
			Name:        "migrate",
			Description: "Run database migration",
			Action: func(c *cli.Context) error {
				err := model.InitDatabase()
				if err != nil {
					log.Println(err)
				}
				err = model.Migrate()
				if err != nil {
					log.Println(err)
				} else {
					log.Println("Migration success")
				}
				return err
			},
		},
		{
			Name:        "seed",
			Description: "Run database seeder",
			Action: func(c *cli.Context) error {
				err := model.InitDatabase()
				if err != nil {
					log.Println(err)
				}
				err = model.RunSeeder()
				if err != nil {
					log.Println("Main.go", err)
				}
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
