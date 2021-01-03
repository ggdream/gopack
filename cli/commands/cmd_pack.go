package commands

import (
	"github.com/ggdream/gopack/cmd/pack"
	"github.com/ggdream/gopack/tools/file"
	"github.com/urfave/cli/v2"
)

func handlePack(c *cli.Context) error {
	t, err := file.JudgeFile()
	if err != nil {
		return err
	}

	return pack.Packs(t, "")
}

// CommPack ...
func CommPack() *cli.Command {

	return &cli.Command{
		Name:   "pack",
		Usage:  "Compile and pack the project (entry `$go tool dist list` to view the supported platforms)",
		Action: handlePack,
	}
}
