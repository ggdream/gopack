package commands

import (
	"errors"
	"github.com/ggdream/gopack/cmd/run"
	"github.com/ggdream/gopack/tools/file"
	"github.com/urfave/cli/v2"
)

func handleRun(c *cli.Context) error {
	t, err := file.JudgeFile()
	if err != nil {
		return err
	}

	if c.NArg() == 0 {
		return errors.New("no command specified")
	}
	return run.Runner(t, "", c.Args().Get(0))
}

// CommRun ...
func CommRun() *cli.Command {

	return &cli.Command{
		Name:   "run",
		Usage:  "Run the script defined in the config file",
		Action: handleRun,
	}
}
