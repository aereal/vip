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
	cli.StringFlag{
		Name:  "lockfile",
		Value: "plugins.lock.json",
		Usage: "Lock file",
	},
}

var Commands = []cli.Command{
	commandInstall,
	commandList,
	commandLock,
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

var commandLock = cli.Command{
	Name:   "lock",
	Usage:  "Output lock file",
	Action: doLock,
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
	env, err := NewEnvironment(c.String("prefix"))
	if err != nil {
		log.Fatal(err)
	}
	for _, deploy := range env {
		println(deploy.Format())
	}
}

func doLock(c *cli.Context) {
	env, err := NewEnvironment(c.String("prefix"))
	if err != nil {
		log.Fatal(err)
	}
	if err = env.Lock(c.String("lockfile")); err != nil {
		log.Fatal(err)
	}
	println("Done")
}

func main() {
	app := cli.NewApp()
	app.Name = "vip"
	app.Version = Version
	app.Author = "aereal"
	app.Commands = Commands

	app.Run(os.Args)
}
