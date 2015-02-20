package main

import (
	"github.com/codegangsta/cli"
	"github.com/davecgh/go-spew/spew"
	// irc "github.com/fluffle/goirc/client"
	"github.com/op/go-logging"
	"os"
	"time"
)

var (
	//jarvis                      *Bot = nil
	//configFilename              = "jarvis.toml"

	log              = logging.MustGetLogger("jarvis")
	logBackendStderr = logging.NewLogBackend(os.Stderr, "", 0)
	logFormat        = logging.MustStringFormatter(
		"%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}",
	)
)

func initLogging() {
	logging.SetFormatter(logFormat)
	logging.SetBackend(logBackendStderr)
	log.Debug("Logging configured")
}

func start(c *cli.Context) {
	log.Debug("Inside start")

	quit := make(chan bool)

	var err error
	var config *JarvisConfig = nil
	var configFile string = c.String("config")

	config, err = loadConfig(configFile)
	if err != nil {
		log.Warning("Unable to load config file %s: %v", configFile, err)
	} else {
		config.filename = configFile
		spew.Dump(config)
	}

	jarvis := newBot(config)
	jarvis.initConnections()
	jarvis.startConnections()

	<-quit

	// connection := &Connection{
	// 	Server:   c.String("server"),
	// 	Port:     c.String("port"),
	// 	Nickname: c.String("nickname"),
	// 	Channels: make(map[string]*ChannelOpts),
	// }

	// connection.Channels[c.String("channel")] = &ChannelOpts{true}

}

func main() {

	// Logging init
	initLogging()

	app := newApp()
	go app.Run(os.Args)

	// jarvis := new(Bot)

	// if config, err := loadConfig(configFilename); err != nil {
	// 	log.Error("Error loading config file %s: %v", configFilename, err)
	// 	os.Exit(1)
	// } else {
	// 	jarvis.Config = config
	// }

	// log.Debug("# of networks defined in config: %d", len(jarvis.Config.Networks))
	// jarvis.initConnections()

	//jarvis.Connections = make(map[string]*Connection)

	// for k, v := range jarvis.Config.Networks {
	// 	// connection, success := jarvis.Connections[k]
	// 	// if success == false {
	// 	// 	connection = v.init()
	// 	// }
	// 	// log.Printf("network [%s], servers = %v", k, v.Servers)
	// 	// jarvis.Connections[k] = connection
	// }

	// Connect to each network
	// for k, v := range jarvis.Connections {
	// 	log.Printf("Connecting to %s", k)
	// 	network, success := jarvis.Config.Networks[k]
	// 	if success {
	// 		v.Connect(network.Servers[0])
	// 	}
	// }

	// For each

	sleepInterval := 60
	for {
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}

}
