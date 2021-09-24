package main

import (
	"fmt"
	"os"

	"github.com/avg73/joker/pkg/actions"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "joker"
	app.Version = "0.0.1"
	app.Usage = "CLI application for https://api.chucknorris.io/"
	app.Commands = []cli.Command{
		{
			Name:   "random",
			Usage:  "prints a random joke.",
			Action: actions.Random,
		},
		{
			Name:  "dump",
			Usage: "uploads n (default: 5) random unique jokes for each of the existing categories and saves them to text files.",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "n",
					Usage: "number of uploaded jokes for each category",
					Value: 5,
				},
			},
			Action: actions.Dump,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
