package commands

import "github.com/urfave/cli/v2"

func SetCommands(app *cli.App) {
	app.Commands = []*cli.Command{
		CommNew(),
		CommInit(),
		CommRun(),
		CommTidy(),
		CommBuild(),
		CommPack(),
		CommNotify(),
	}
}
