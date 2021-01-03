package commands

import (
	"fmt"
	"strings"

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
	var repo string
	dic := c.Args().Get(0) // dic即文件名，不带前缀
	if len(strings.Split(dic, "/")) == 3 {
		repo = dic
	} else {
		repo = fmt.Sprintf("github.com/%s/%s", name, dic)
	}

	return new2.Create(global.ConfigFileType, "", dic, dic, global.DefaultVersion,
		fmt.Sprintf("<%s %s>", name, email), global.DefaultDesc, repo)
}

// CommNew ...
func CommNew() *cli.Command {

	return &cli.Command{
		Name:   "new",
		Usage:  "Create new project",
		Flags:  []cli.Flag{FlagType},
		Action: handleNew,
	}
}
