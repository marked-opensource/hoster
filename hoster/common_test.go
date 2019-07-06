package hoster_test

import (
	"flag"
	"github.com/urfave/cli"
	"testing"
)

func fakeCliContext(t *testing.T, args []string) *cli.Context {
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	err := set.Parse(args)
	if err != nil {
		t.Log("Could not parse test args")
		t.Fail()
	}
	return cli.NewContext(app, set, nil)
}
