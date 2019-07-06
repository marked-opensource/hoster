package hoster_test

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
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

func captureOutput(t *testing.T, f func()) string {
	rescueStdout := os.Stdout
	r, w, err := os.Pipe()
	assert.NoError(t, err)
	os.Stdout = w
	f()
	err = w.Close()
	assert.NoError(t, err)
	out, err := ioutil.ReadAll(r)
	assert.NoError(t, err)
	os.Stdout = rescueStdout
	return string(out)
}
