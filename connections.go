package main

import (
	_ "github.com/davecgh/go-spew/spew"
	"github.com/nickvanw/ircx"
	_ "github.com/sorcix/irc"
	"net"
)

type Connection struct {
	bot    *ircx.Bot
	config *Network
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
		_, success := j.connections[k]
		if !success {
			j.connections[k] = v.newConnection()
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
