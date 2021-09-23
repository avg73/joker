package actions

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"
)

func Dump(c *cli.Context) error {
	n := c.Int("n")
	fmt.Println("Dump jokes", n)
	return nil
}
