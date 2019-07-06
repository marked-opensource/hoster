package main

import (
	"hoster/hoster"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "hoster"
	app.Usage = "env manager for /etc/hosts"
	app.Version = os.Getenv("HOSTER_BUILD_VER")

	app.Commands = []cli.Command{
		{
			Name:    "activate",
			Aliases: []string{"a"},
			Usage:   "activate given host env",
			Action:  hoster.NewActivationHandler().Handle,
		},
		{
			Name:    "deactivate",
			Aliases: []string{"a"},
			Usage:   "deactivate given host env",
			Action:  hoster.NewDeactivationHandler().Handle,
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
