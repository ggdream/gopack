package commands

import "github.com/urfave/cli/v2"


// SetCommands ...
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
