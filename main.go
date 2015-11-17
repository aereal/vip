package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var Version string

var Commands = []cli.Command{
}

func main() {
	app := cli.NewApp()
	app.Name = "vip"
	app.Version = Version
	app.Author = "aereal"
	app.Commands = Commands

	app.Run(os.Args)
}

