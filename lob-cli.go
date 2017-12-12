package main

import (
	"fmt"
	"os"

	"github.com/davidnix/lob-cli/postcard"
	dotenv "github.com/joho/godotenv"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "lob-cli"
	app.Usage = "Print postcards and such from lob.com"
	app.Version = "0.0.2"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "api-key, k",
			Usage:  "Lob.com API `KEY`",
			EnvVar: "LOB_API_KEY",
		},
		cli.StringFlag{
			Name:  "csv",
			Usage: "Load recipient addresses from csv `FILE`",
		},
		cli.StringFlag{
			Name:   "from-name",
			Usage:  "Return address `NAME`",
			EnvVar: "LOB_FROM_NAME",
		},
		cli.StringFlag{
			Name:   "from-address",
			Usage:  "Return address `STREET`",
			EnvVar: "LOB_FROM_ADDRESS",
		},
		cli.StringFlag{
			Name:   "from-city",
			Usage:  "Return address `CITY`",
			EnvVar: "LOB_FROM_CITY",
		},
		cli.StringFlag{
			Name:   "from-state",
			Usage:  "Return address `STATE`",
			EnvVar: "LOB_FROM_STATE",
		},
		cli.StringFlag{
			Name:   "from-zip",
			Usage:  "Return address `ZIPCODE`",
			EnvVar: "LOB_FROM_ZIPCODE",
		},
		cli.StringFlag{
			Name:   "from-country",
			Usage:  "Return address `COUNTRY`",
			EnvVar: "LOB_FROM_COUNTRY",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "postcards",
			Usage:  "Send postcards",
			Action: postcard.Send,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "front",
					Usage: "Load front from html template `FILE`",
				},
				cli.StringFlag{
					Name:  "back",
					Usage: "Load back from html template `FILE`",
				},
			},
		},
	}

	loadEnv()

	if runErr := app.Run(os.Args); runErr != nil {
		fmt.Println("RUN ERROR:", runErr)
	}
}

func loadEnv() {
	// file exists
	if _, err := os.Stat("./.env"); err == nil {
		if envErr := dotenv.Load(); envErr != nil {
			fmt.Println("Warning, unable to load .env:", envErr)
		}
	}
}
