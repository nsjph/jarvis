package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Jarvis"
	app.Usage = "A helpful IRC bot"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: fmt.Sprintf("%s/.jarvis.toml", os.Getenv("HOME")),
			Usage: "Config file location",
		},
		cli.StringFlag{
			Name:  "server, s",
			Value: "irc.freenode.net",
			Usage: "IRC server to connect to",
		},
		cli.StringFlag{
			Name:  "port, p",
			Value: "6667",
			Usage: "IRC port to connect to",
		},
		cli.StringFlag{
			Name:  "nickname, n",
			Value: "Jarvis",
			Usage: "Nickname to use",
		},
		cli.StringFlag{
			Name:  "channel",
			Value: "#jarvis-test",
			Usage: "Default IRC channel to join on connect",
		},
	}

	app.Action = start

	return app
}
