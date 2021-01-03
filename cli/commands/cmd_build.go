package commands

import (
	"github.com/ggdream/gopack/cmd/build"
	"github.com/ggdream/gopack/tools/file"
	"github.com/urfave/cli/v2"
)

func handleBuild(c *cli.Context) error {
	t, err := file.JudgeFile()
	if err != nil {
		return err
	}

	return build.Builds(t, "")
}

// CommBuild ...
func CommBuild() *cli.Command {

	return &cli.Command{
		Name:   "build",
		Usage:  "Compile the project (entry `$go tool dist list` to view the supported platforms)",
		Action: handleBuild,
	}
}
