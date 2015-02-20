package main

import (
	"github.com/codegangsta/cli"
	_ "github.com/davecgh/go-spew/spew"
	_ "github.com/nickvanw/ircx"
)

type Jarvis struct {
	connections map[string]*Connection
	config      *JarvisConfig
}

func start(c *cli.Context) {
	log.Debug("Inside start")

	var err error
	var config *JarvisConfig = nil
	var configFile string = c.String("config")

	config, err = loadConfig(configFile)
	if err != nil {
		log.Warning("Unable to load config file %s: %v", configFile, err)
	} else {
		config.filename = configFile
	}

	jarvis := &Jarvis{
		connections: make(map[string]*Connection),
		config:      config,
	}

	jarvis.initConnections()
	for k, v := range jarvis.connections {
		err := v.startConnection()
		if err != nil {
			log.Error("Error connecting to %s: %v", k, err)
		}
	}
}
