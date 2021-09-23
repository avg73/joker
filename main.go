package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "joker"
	app.Usage = "CLI application for https://api.chucknorris.io/"
	app.Commands = []cli.Command{
		{
			Name:   "random",
			Usage:  "prints a random joke.",
			Action: randomJoke,
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
			Action: dumpJokes,
		},
	}
	app.Run(os.Args)
}

func randomJoke(c *cli.Context) error {
	fmt.Println("Random joke")
	return nil
}

func dumpJokes(c *cli.Context) error {
	n := c.Int("n")
	fmt.Println("Dump jokes", n)
	return nil
}
