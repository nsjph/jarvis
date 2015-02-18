package main

import (
	"crypto/tls"
	"github.com/BurntSushi/toml"
	_ "github.com/davecgh/go-spew/spew"
	irc "github.com/thoj/go-ircevent"
	"log"
	_ "regexp"
	"strings"
	"time"
)

type Bot struct {
	Connections map[string]*irc.Connection
	Config      *JarvisConfig
	//Plugins     map[string]*Plugin
}

type JarvisConfig struct {
	Networks map[string]Network
}

type Network struct {
	Nickname string
	Realname string
	Version  string
	Altnames []string
	Servers  []string
	Channels []string
}

var (
	jarvis                      *Bot = nil
	tlsEnabled                       = true
	tlsVerify                        = true
	tlsVersion                       = tls.VersionTLS12
	tlsPreferServerCipherSuites      = false
	tlsCipherSuites                  = []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256}
)

func parseConfig(f string) (*JarvisConfig, error) {

	config := new(JarvisConfig)

	if _, err := toml.DecodeFile(f, config); err != nil {
		return nil, err
	}

	return config, nil
}

func newTLSConfig(serverName string) *tls.Config {
	log.Printf("newTLSConfig for %s", serverName)
	config := new(tls.Config)
	config.ServerName = serverName
	config.InsecureSkipVerify = tlsVerify
	//config.CipherSuites = tlsCipherSuites
	config.PreferServerCipherSuites = tlsPreferServerCipherSuites

	return config

}

func addCallbacks(c *irc.Connection) {
	log.Println("addCallbacks")
	c.AddCallback("001", logEvent)
	c.AddCallback("PRIVMSG", logEvent)
}

func logEvent(event *irc.Event) {
	log.Printf("[%s] (%s) %s", event.Host, event.Nick, event.Message)
}

func (n Network) init() *irc.Connection {
	connection := irc.IRC(n.Nickname, n.Realname)
	connection.VerboseCallbackHandler = true
	connection.Debug = true
	hostport := strings.Split(n.Servers[0], ":")
	if len(hostport) != 2 {
		log.Fatalf("Server host:port (%s) not specified correctly", n.Servers[0])
	}

	host := hostport[0]

	if tlsEnabled == true {
		connection.UseTLS = true
		connection.TLSConfig = newTLSConfig(host)
	}

	addCallbacks(connection)

	return connection
}

func main() {

	config, err := parseConfig("jarvis.toml")
	if err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	jarvis := new(Bot)
	jarvis.Config = config
	jarvis.Connections = make(map[string]*irc.Connection)

	for k, v := range jarvis.Config.Networks {
		connection, success := jarvis.Connections[k]
		if success == false {
			connection = v.init()
		}
		log.Printf("network [%s], servers = %v", k, v.Servers)
		jarvis.Connections[k] = connection
	}

	// Connect to each network
	for k, v := range jarvis.Connections {
		log.Printf("Connecting to %s", k)
		network, success := jarvis.Config.Networks[k]
		if success {
			v.Connect(network.Servers[0])
		}
	}

	// For each

	sleepInterval := 60
	for {
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}

}
