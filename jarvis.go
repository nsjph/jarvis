package main

import (
	"github.com/boltdb/bolt"
	"github.com/codegangsta/cli"
	//cli "github.com/jawher/mow.cli"
)

type Jarvis struct {
	connections map[string]*Connection
	config      *JarvisConfig
	db          *bolt.DB
}

func start(c *cli.Context) {
	var err error

	// log.Debug("Inside start")
	// plugins, err := pluginsAvailable(os.Getenv("HOME") + "/.jarvis/plugins/")
	// if err != nil {
	// 	log.Error("Failed to check for plugins: %v", err)
	// 	return
	// }
	// loadGlispPlugins(plugins)

	// return

	var config *JarvisConfig = nil
	var configFile string = c.String("config")

	config, err = loadConfig(configFile)
	if err != nil {
		log.Warning("Unable to load config file %s: %v", configFile, err)
	} else {
		config.filename = configFile
	}

	config.pluginDirectory = c.String("plugins")

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
