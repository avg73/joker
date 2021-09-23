package actions

import (
	"fmt"
	"os"

	"github.com/avg73/joker/pkg/repository"
	"golang.org/x/sync/errgroup"
	"gopkg.in/urfave/cli.v1"
)

func Dump(c *cli.Context) error {
	n := c.Int("n")

	categories, err := repository.GetCategories()
	if err != nil {
		return err
	}

	g := new(errgroup.Group)
	for _, c := range categories {
		category := c
		g.Go(func() error { return getAndWriteJokes(n, category) })
	}
	return g.Wait()
}

func getAndWriteJokes(n int, category string) error {
	jokes := make([]string, n)

	for i := range jokes {
		joke, err := repository.GetJoke(category)
		if err != nil {
			return err
		}
		jokes[i] = joke.Value
	}

	name := fmt.Sprintf("%s.txt", category)
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	for _, j := range jokes {
		fmt.Fprintln(f, j)
	}
	return f.Close()
}
