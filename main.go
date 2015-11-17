package main

import (
	"log"
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
	cli.StringFlag{
		Name:  "prefix",
		Value: os.ExpandEnv("$HOME/.vim/bundle"),
		Usage: "Custom prefix of installation path",
	},
}

var Commands = []cli.Command{
	commandInstall,
	commandList,
}

var commandInstall = cli.Command{
	Name:   "install",
	Usage:  "Install plugins",
	Action: doInstall,
	Flags:  CommonFlags,
}

var commandList = cli.Command{
	Name:   "list",
	Usage:  "List installed plugins",
	Action: doList,
	Flags:  CommonFlags,
}

func doInstall(c *cli.Context) {
	recipe, err := NewRecipeFromManifestJSON(c.String("manifest"))
	if err != nil {
		log.Fatal(err)
	}
	BatchInstall(recipe, c.String("prefix"))
}

func doList(c *cli.Context) {
	index, err := NewLocalRepositoryIndexFromPrefix(c.String("prefix"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", index)
}

func main() {
	app := cli.NewApp()
	app.Name = "vip"
	app.Version = Version
	app.Author = "aereal"
	app.Commands = Commands

	app.Run(os.Args)
}
