package main

import (
	_ "github.com/davecgh/go-spew/spew"
	"github.com/nickvanw/ircx"
	"github.com/sorcix/irc"
)

func (c *Connection) RegisterHandlers() {
	c.bot.AddCallback(irc.RPL_WELCOME, ircx.Callback{Handler: ircx.HandlerFunc(c.RegisterConnect)})
	c.bot.AddCallback(irc.PING, ircx.Callback{Handler: ircx.HandlerFunc(PingHandler)})
}

func (c *Connection) RegisterConnect(s ircx.Sender, m *irc.Message) {
	log.Debug("RegisterConnect")

	for _, v := range c.config.Channels {
		log.Debug("Joining %s", v)
		s.Send(&irc.Message{
			Command: irc.JOIN,
			Params:  []string{v},
		})
	}
}

func PingHandler(s ircx.Sender, m *irc.Message) {
	s.Send(&irc.Message{
		Command:  irc.PONG,
		Params:   m.Params,
		Trailing: m.Trailing,
	})
}
