package main

import (
	"github.com/nickvanw/ircx"
	"github.com/sorcix/irc"
)

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

func (c *Connection) privmsgHandler(s ircx.Sender, m *irc.Message) {
	log.Debug("cmd [%s], params [%s], trailing [%s]", m.Command, m.Params, m.Trailing, m.EmptyTrailing)

	if startsWithPrefix(m.Trailing) {
		command, arguments := parseCommand(m.Trailing)
		target := m.Params[0]
		if command != "" {
			log.Debug("received command [%s], args [%v], target %s", command, arguments, target)
		}

	}

}
