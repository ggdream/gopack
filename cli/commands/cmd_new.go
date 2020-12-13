package commands

import (
	"fmt"
	new2 "github.com/ggdream/gopack/cmd/new"
	"github.com/ggdream/gopack/global"
	"github.com/ggdream/gopack/tools/git"
	"github.com/urfave/cli/v2"
)

func handleNew(c *cli.Context) error {
	name, email, err := git.GetUserInfo()
	if err != nil {
		return err
	}
	dic := c.Args().Get(0) // dic即文件名，不带前缀

	return new2.Create(global.ConfigFileType, "", dic, dic, global.DefaultVersion,
		fmt.Sprintf("<%s %s>", name, email), global.DefaultDesc, fmt.Sprintf("github.com/%s/%s", name, dic))
}

func CommNew() *cli.Command {

	return &cli.Command{
		Name:   "new",
		Usage:  "Create new project",
		Flags:  []cli.Flag{FlagType},
		Action: handleNew,
	}
}
