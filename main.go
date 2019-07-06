package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "hoster"
	app.Usage = "env manager for /etc/hosts"

	app.Commands = []cli.Command{
		{
			Name:    "activate",
			Aliases: []string{"a"},
			Usage:   "activate given host env",
			Action: func(c *cli.Context) {
				fmt.Printf("TODO activate %s\n", c.Args().First())
			},
		},
		{
			Name:    "deactivate",
			Aliases: []string{"a"},
			Usage:   "deactivate given host env",
			Action: func(c *cli.Context) {
				fmt.Printf("TODO deactivate %s\n", c.Args().First())
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
