package commands

import (
	"github.com/ggdream/gopack/cmd/tidy"
	"github.com/ggdream/gopack/tools/file"
	"github.com/urfave/cli/v2"
)

func handleTidy(c *cli.Context) error {
	t, err := file.JudgeFile()
	if err != nil {
		return err
	}

	return tidy.Tidy(t, "")
}

// CommTidy ...
func CommTidy() *cli.Command {

	return &cli.Command{
		Name:   "tidy",
		Usage:  "Tidy all dependencies about the project",
		Action: handleTidy,
	}
}
