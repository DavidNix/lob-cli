package main

import (
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "lob-cli"
	app.Usage = "Print postcards and such from lob.com"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "api-key, k",
			Usage:  "Lob.com API `KEY`",
			EnvVar: "LOB_API_KEY",
		},
		cli.StringFlag{
			Name:  "csv",
			Usage: "Load addresses from `FILE`",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "postcards",
			Usage: "Send postcards",
			Action: func(c *cli.Context) error {
				fmt.Println("postcard!")
				return nil
			},
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

	if err := app.Run(os.Args); err != nil {
		fmt.Println("ERROR:", err)
	}
}
