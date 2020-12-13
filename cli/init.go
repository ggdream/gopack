package cli

import (
	"github.com/ggdream/gopack/cli/commands"
	"github.com/ggdream/gopack/global"
	"github.com/urfave/cli/v2"
	"os"
)

func New() error {
	app := cli.NewApp()

	app.Name = global.NAME
	app.Version = global.VERSION
	app.Usage = global.USAGE
	app.Authors = func() []*cli.Author {
		authors := make([]*cli.Author, 0)
		for k, v := range global.AUTHORS {
			authors = append(authors, &cli.Author{Name: k, Email: v})
		}
		return authors
	}()

	commands.SetCommands(app)
	return app.Run(os.Args)
}
