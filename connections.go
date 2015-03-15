package main

import (
	"github.com/boltdb/bolt"
	"github.com/nickvanw/ircx"
	"net"
)

type Connection struct {
	name    string
	bot     *ircx.Bot
	config  *Network
	plugins []*GlispPlugin
	db      *bolt.DB
}

func (n *Network) newConnection() *Connection {

	var transport string = ""

	if n.UseIPv6 == true {
		transport = "tcp6"
	} else {
		transport = "tcp"
	}

	host, _, err := net.SplitHostPort(n.Servers[0])
	if err != nil {
		log.Fatalf("Error parsing %s into host:port", n.Servers[0])
	}

	tcpaddr, err := net.ResolveTCPAddr(transport, n.Servers[0])
	if err != nil {
		log.Fatalf("newConnection: Failed looking up %s: %v", host, err)
	} else {
		log.Debug("Resolved server to %s", tcpaddr.String())
	}

	bot := ircx.Classic(tcpaddr.String(), n.Nickname)
	//plugins, err := initPlugins()

	return &Connection{
		bot:    bot,
		config: n,
	}
}

func (j *Jarvis) initConnections() (err error) {

	if j.connections == nil {
		log.Warning("initConnections: initializing connections map")
		j.connections = make(map[string]*Connection)
	}

	for k, v := range j.config.Networks {
		connection, success := j.connections[k]
		if !success {
			connection = v.newConnection()
			connection.db = j.db
			connection.name = k
			// plugins, err := initPlugins(j.config.pluginDirectory)
			// if err != nil {
			// 	log.Error("Unable to init plugins: %v", err)
			// } else {
			// 	connection.plugins = plugins
			// }
			j.connections[k] = connection
		}
	}

	return nil
}

func (c *Connection) startConnection() error {
	log.Debug("starting connection to %s", c.bot.Server)

	err := c.bot.Connect()
	if err != nil {
		log.Debug("Error connecting to %s: %v", c.bot.Server, err)
		return err
	}

	c.RegisterHandlers()

	c.bot.CallbackLoop()

	return nil
}

// type Callback struct {
// 	Handler Handler
// 	Sender  Sender
// }

// type Handler interface {
// 	Handle(Sender, *irc.Message)
// }

// type Sender interface {
// 	Send(*irc.Message) error
// }

// type HandlerFunc func(s Sender, m *irc.Message)

// func (f HandlerFunc) Handle(s Sender, m *irc.Message) {
// 	f(s, m)
// }

// // Implement ircx.Handler interface

// func (c *Connection) Handle(s Sender, m *irc.Message) {
// 	log.Debug("Handle")
// 	spew.Dump(s)
// 	spew.Dump(m)
// }

// Implement ircx.Sender interface

// func (c *Connection) Send(m *irc.Message) error {
// 	log.Debug("Send")
// 	spew.Dump(m)
// 	return nil
// }
