package main

import (
	"github.com/nickvanw/ircx"
	"github.com/sorcix/irc"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var (
	prefix  = "@"
	plugins = make(map[string]func(*Connection, ircx.Sender, *irc.Message, string, []string))
)

func (c *Connection) RegisterHandlers() {
	c.bot.AddCallback(irc.RPL_WELCOME, ircx.Callback{Handler: ircx.HandlerFunc(c.RegisterConnect)})
	c.bot.AddCallback(irc.PRIVMSG, ircx.Callback{Handler: ircx.HandlerFunc(c.privmsgHandler)})
	c.bot.AddCallback(irc.PING, ircx.Callback{Handler: ircx.HandlerFunc(PingHandler)})
}

func startsWithPrefix(s string) bool {
	if string(s[0]) == prefix {
		return true
	}
	return false
}

func parseCommand(s string) (command string, arguments []string) {
	params := strings.Fields(strings.Split(s, prefix)[1])

	if len(params) == 0 {
		return "", nil
	}

	command = params[0]
	arguments = params[1:]
	return command, arguments
}

// func pluginsAvailable(pluginDirectory string) ([]string, error) {
// 	entries, err := ioutil.ReadDir(pluginDirectory)
// 	if err != nil {
// 		log.Debug("Unable to read plugin directory (%s): %v", pluginDirectory, err)
// 		return nil, err
// 	}

// 	var plugins []string

// 	for _, v := range entries {
// 		if v.IsDir() == false {
// 			ext := path.Ext(v.Name())
// 			if ext == ".glisp" {
// 				plugins = append(plugins, path.Join(pluginDirectory, v.Name()))
// 			}
// 		}
// 	}
// 	return plugins, nil
// }

// func initPlugins(dir string) ([]*GlispPlugin, error) {

// 	log.Debug("Finding available plugins")
// 	available, err := pluginsAvailable(dir)
// 	if err != nil {
// 		return nil, err
// 	}

// 	log.Debug("Loading available plugins")
// 	plugins, err := loadGlispPlugins(available)
// 	if err != nil {
// 		return nil, err
// 	}

// 	os.Exit(0)

// 	return plugins, nil

// }

/////////// Basic golang-native plugins below
