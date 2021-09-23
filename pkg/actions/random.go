package actions

import (
	"fmt"

	"github.com/avg73/joker/pkg/repository"
	"gopkg.in/urfave/cli.v1"
)

func Random(c *cli.Context) error {
	joke, err := repository.GetJoke("")
	if err != nil {
		return err
	}
	fmt.Println(joke.Value)
	return nil
}
