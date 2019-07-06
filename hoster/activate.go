package hoster

import (
	"fmt"
	"github.com/urfave/cli"
)

type activationHandler struct{}

func (h *activationHandler) Handle(c *cli.Context) {
	fmt.Printf("TODO activate %s\n", c.Args().First())
}

func NewActivationHandler() CommandHandler {
	return &activationHandler{}
}
