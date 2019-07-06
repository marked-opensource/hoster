package hoster

import (
	"fmt"
	"github.com/urfave/cli"
)

type deactivationHandler struct{}

func (h *deactivationHandler) Handle(c *cli.Context) {
	fmt.Printf("TODO deactivate %s\n", c.Args().First())
}

func NewDeactivationHandler() CommandHandler {
	return &deactivationHandler{}
}
