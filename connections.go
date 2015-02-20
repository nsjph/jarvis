package main

import (
	irc "github.com/nsjph/goirc/client"
)

// type Connection struct {
// 	*client.Conn
// 	network *Network
// }

type ChannelOpt struct {
	autojoin bool
}

func (b *Bot) startConnections() error {
	for network, conn := range b.connections {
		log.Info("Connecting to %s (%s)", network, conn.Config().Server)
		if err := conn.Connect(); err != nil {
			log.Error("Connection error: ", err)
			return err
		} else {
			log.Info("Connected to %s", network)
			//conn.Join("jarvis-test")
		}
	}
	return nil
}

func newConnection(n *Network) *irc.Conn {

	cfg := irc.NewConfig(n.Nickname)
	cfg.Server = n.Servers[0]
	cfg.Version = n.Version
	cfg.QuitMessage = "It's full of stars..."
	//cfg.Me.Ident = "jarvis"
	//cfg.Me.Name = n.Realname
	cfg.UseIPv6 = n.UseIPv6
	cfg.NewNick = func(n string) string { return n + "_" }

	conn := irc.Client(cfg)

	conn.HandleFunc("CONNECTED", func(conn *irc.Conn, line *irc.Line, config *Network) {
		log.Debug("CONNECTED on network %s", n.Nickname)
	})
	conn.HandleFunc("connected",
		func(conn *irc.Conn, line *irc.Line) {
			log.Debug("connected")
			conn.Join("#jarvis-test")
		})

	return conn

	//return &Connection{conn, n}
}

func (b *Bot) initConnections() {

	if b.config == nil {
		log.Warning("initConnections: No connections to initialize, config empty")
		return
	}

	if b.connections == nil {
		log.Debug("initConnections: Initializing bot.connections map")
		b.connections = make(map[string]*irc.Conn)
	}

	for k, v := range b.config.Networks {
		conn, success := b.connections[k]
		if success == false {
			log.Debug("initConnections: Creating connection for [%s]", k)
			conn = newConnection(&v)
			//conn.addDefaultHandlers()
			b.connections[k] = conn
		} else {
			log.Debug("initConnections: Connection exists for [%s]", k)
		}
	}
}
