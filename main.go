package main

import (
	"crypto/tls"
	"github.com/BurntSushi/toml"
	"github.com/op/go-logging"
	"github.com/sorcix/irc"
	"net"
	"os"
	_ "strings"
	"time"
)

type Connection struct {
	Server    string
	Port      string
	UseTLS    bool
	TLSConfig *tls.Config
	Nickname  string
	Realname  string
	Version   string
	Data      chan *irc.Message
	Reader    *irc.Decoder
	Writer    *irc.Encoder
	Conn      *irc.Conn
	tlsConn   *tls.Conn
	tcpConn   *net.Conn
}

type Bot struct {
	Connections map[string]*Connection
	Config      *JarvisConfig
}

type JarvisConfig struct {
	Networks map[string]Network
}

type Network struct {
	Nickname  string
	Realname  string
	Version   string
	UseTLS    bool `toml:"useTLS"`
	VerifyTLS bool `toml:"verifyTLS"`
	Altnames  []string
	Servers   []string
	Channels  []string
}

var (
	jarvis                      *Bot = nil
	configFilename                   = "jarvis.toml"
	tlsEnabled                       = true
	tlsVerify                        = true
	tlsVersion                       = tls.VersionTLS12
	tlsPreferServerCipherSuites      = false
	tlsCipherSuites                  = []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256}
	log              = logging.MustGetLogger("jarvis")
	logBackendStderr = logging.NewLogBackend(os.Stderr, "", 0)
	logFormat        = logging.MustStringFormatter(
		"%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}",
	)
)

func newTLSConfig(server string, verify bool) *tls.Config {
	log.Debug("newTLSConfig for %s", server)
	config := new(tls.Config)
	config.ServerName = server
	config.InsecureSkipVerify = verify
	config.PreferServerCipherSuites = tlsPreferServerCipherSuites
	//config.CipherSuites = tlsCipherSuites

	return config
}

func loadConfig(filename string) (*JarvisConfig, error) {

	config := new(JarvisConfig)

	if _, err := toml.DecodeFile(filename, config); err != nil {
		return nil, err
	}

	return config, nil
}

func initLogging() {
	logging.SetFormatter(logFormat)
	logging.SetBackend(logBackendStderr)
	log.Debug("Logging configured")
}

func main() {

	// Logging init
	initLogging()

	jarvis := new(Bot)

	if config, err := loadConfig(configFilename); err != nil {
		log.Error("Error loading config file %s: %v", configFilename, err)
		os.Exit(1)
	} else {
		jarvis.Config = config
	}

	log.Debug("# of networks defined in config: %d", len(jarvis.Config.Networks))
	jarvis.initConnections()

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
