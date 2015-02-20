package main

import (
	irc "github.com/nsjph/goirc/client"
)

type Bot struct {
	connections map[string]*irc.Conn
	//connections map[string]*Connection
	config *JarvisConfig
	//log         map[string]*logging.Logger
}

func newBot(config *JarvisConfig) *Bot {

	b := new(Bot)
	b.connections = make(map[string]*irc.Conn)
	b.config = config

	return b
}

// func newBotOld(config *JarvisConfig) *Bot {

// 	b := new(Bot)
// 	b.connections = make(map[string]*irc.Conn)
// 	b.config = config

// 	return b
// }
