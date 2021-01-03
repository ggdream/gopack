package commands

import (
	"github.com/ggdream/gopack/cmd/notify"
	"github.com/urfave/cli/v2"
)

func handleNotify(c *cli.Context) error {
	notify.Notify()
	return nil
}

// CommNotify ...
func CommNotify() *cli.Command {

	return &cli.Command{
		Name:   "notify",
		Usage:  "Tell you what i want to tell you",
		Action: handleNotify,
	}
}
