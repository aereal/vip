package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var Version string

var CommonFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "manifest",
		Value: "plugins.json",
		Usage: "",
	},
}

var Commands = []cli.Command{
	commandInstall,
}

var commandInstall = cli.Command{
	Name:   "install",
	Usage:  "Install plugins",
	Action: doInstall,
	Flags:  CommonFlags,
}

func doInstall(c *cli.Context) {
	recipe := NewRecipeFromManifestJSON(c.String("manifest"))
	BatchInstall(recipe)
}

func main() {
	app := cli.NewApp()
	app.Name = "vip"
	app.Version = Version
	app.Author = "aereal"
	app.Commands = Commands

	app.Run(os.Args)
}
