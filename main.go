package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var Version string

var pluginsFile string

var Commands = []cli.Command{
	commandInstall,
}

var commandInstall = cli.Command{
	Name:   "install",
	Usage:  "Install plugins",
	Action: doInstall,
}

func doInstall(c *cli.Context) {
	recipe := NewRecipeFromManifestJSON(pluginsFile)
	BatchInstall(recipe)
}

func main() {
	app := cli.NewApp()
	app.Name = "vip"
	app.Version = Version
	app.Author = "aereal"
	app.Commands = Commands

	pluginsFile = "plugins.json" // TODO

	app.Run(os.Args)
}
