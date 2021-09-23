package actions

import (
	"fmt"
	"os"

	"github.com/avg73/joker/pkg/repository"
	"gopkg.in/urfave/cli.v1"
)

func Dump(c *cli.Context) error {
	n := c.Int("n")

	categories, err := repository.GetCategories()
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, c := range categories {
		jokes := make([]string, n)

		for i := range jokes {
			joke, err := repository.GetJoke(c)
			if err != nil {
				fmt.Println(err)
				return err
			}
			jokes[i] = joke.Value
		}

		err = writeFile(c, jokes)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeFile(name string, jokes []string) error {
	name += ".txt"
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	for _, j := range jokes {
		fmt.Fprintln(f, j)
	}
	return f.Close()
}
