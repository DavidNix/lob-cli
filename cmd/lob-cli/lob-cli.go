package main

import (
	"fmt"
	"os"

	"github.com/DavidNix/lob-cli/postcard"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "lob-cli"
	app.Usage = "Print postcards and such from lob.com"
	app.Version = "0.0.3"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "api-key, k",
			Usage:   "Lob.com API `KEY`",
			EnvVars: []string{"LOB_API_KEY"},
		},
		&cli.StringFlag{
			Name:  "csv",
			Usage: "Load recipient addresses from csv `FILE`",
		},
		&cli.StringFlag{
			Name:    "from-name",
			Usage:   "Return address `NAME`",
			EnvVars: []string{"LOB_FROM_NAME"},
		},
		&cli.StringFlag{
			Name:    "from-address",
			Usage:   "Return address `STREET`",
			EnvVars: []string{"LOB_FROM_ADDRESS"},
		},
		&cli.StringFlag{
			Name:    "from-city",
			Usage:   "Return address `CITY`",
			EnvVars: []string{"LOB_FROM_CITY"},
		},
		&cli.StringFlag{
			Name:    "from-state",
			Usage:   "Return address `STATE`",
			EnvVars: []string{"LOB_FROM_STATE"},
		},
		&cli.StringFlag{
			Name:    "from-zip",
			Usage:   "Return address `ZIPCODE`",
			EnvVars: []string{"LOB_FROM_ZIPCODE"},
		},
		&cli.StringFlag{
			Name:    "from-country",
			Usage:   "Return address `COUNTRY`",
			EnvVars: []string{"LOB_FROM_COUNTRY"},
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:   "postcards",
			Usage:  "Send postcards",
			Action: postcard.Send,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "front",
					Usage: "Load front from html template `FILE`",
				},
				&cli.StringFlag{
					Name:  "back",
					Usage: "Load back from html template `FILE`",
				},
			},
		},
	}

	if runErr := app.Run(os.Args); runErr != nil {
		fmt.Println("RUN ERROR:", runErr)
	}
}
