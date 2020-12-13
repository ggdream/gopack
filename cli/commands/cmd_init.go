package commands

import (
	"fmt"
	Init "github.com/ggdream/gopack/cmd/init"
	"github.com/ggdream/gopack/global"
	"github.com/ggdream/gopack/tools/git"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
)

var (
	FlagType = &cli.StringFlag{
		Name:        "type",
		Aliases:     []string{"t"},
		Value:       global.ConfigFileType,
		Destination: &global.ConfigFileType,
		Usage:       "Choose any format in `json`, `yaml`, `toml`",
	}
)

func handleInit(c *cli.Context) error {
	name, _, err := git.GetUserInfo()
	if err != nil {
		return err
	}

	dir, _ := os.Getwd()
	_, dic := filepath.Split(dir)
	return Init.Generate(global.ConfigFileType, "", dic, global.DefaultVersion,
		name, global.DefaultDesc, fmt.Sprintf("github.com/%s/%s", name, dic))
}

func CommInit() *cli.Command {

	return &cli.Command{
		Name:   "init",
		Usage:  "Generate gopack configuration file in existing project folder",
		Flags:  []cli.Flag{FlagType},
		Action: handleInit,
	}
}
