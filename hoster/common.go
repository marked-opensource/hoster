package hoster

import "github.com/urfave/cli"

type CommandHandler interface {
	Handle(contest *cli.Context)
}
