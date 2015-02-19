package main

import (
	"crypto/tls"
	_ "github.com/op/go-logging"
	"github.com/sorcix/irc"
	"net"
)

func newConnection(n Network) (*Connection, error) {
	var err error
	c := new(Connection)
	c.Server, c.Port, err = net.SplitHostPort(n.Servers[0])
	if err != nil {
		log.Debug("Error parsing %s: %v", n.Servers[0], err)
		return nil, err
	}

	if n.UseTLS == true {
		c.UseTLS = true
		c.TLSConfig = newTLSConfig(c.Server, n.VerifyTLS)
		c.tlsConn, err = tls.Dial("tcp", c.Server+":"+c.Port, c.TLSConfig)
		if err != nil {
			log.Fatalf("Unable to establish TLS connection to %s: %v", c.Server, err)
		} else {
			log.Debug("Connected to %s", c.Server)
			c.Conn = irc.NewConn(c.tlsConn)
			c.Reader = irc.NewDecoder(c.tlsConn)
			c.Writer = irc.NewEncoder(c.tlsConn)
		}
	}
	return c, nil

}

func (b *Bot) initConnections() {
	var err error

	if b.Connections == nil {
		log.Debug("Initializing new Connections struct")
		b.Connections = make(map[string]*Connection)
	}

	for k, v := range b.Config.Networks {
		connection, success := b.Connections[k]
		if success == false {
			connection, err = newConnection(v)
			if err != nil {
				log.Fatalf("Failed to create connection for network [%s]: %v", k, err)
			}
			b.Connections[k] = connection
		}
	}
}
