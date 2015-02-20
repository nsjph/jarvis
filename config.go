package main

import (
	"crypto/tls"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
	"os"
)

var (
	tlsEnabled                  = true
	tlsVerify                   = true
	tlsVersion                  = tls.VersionTLS12
	tlsPreferServerCipherSuites = false
	tlsCipherSuites             = []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256}
)

type JarvisConfig struct {
	filename string
	Networks map[string]Network
}

type Network struct {
	Nickname  string
	Realname  string
	Version   string
	UseTLS    bool `toml:"useTLS"`
	VerifyTLS bool `toml:"verifyTLS"`
	UseIPv6   bool `toml:"ipv6"`
	Altnames  []string
	Servers   []string
	Channels  []string
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Jarvis"
	app.Usage = "A helpful IRC bot"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: fmt.Sprintf("%s/.jarvis.toml", os.Getenv("HOME")),
			Usage: "Config file location",
		},
		cli.StringFlag{
			Name:  "server, s",
			Value: "irc.freenode.net",
			Usage: "IRC server to connect to",
		},
		cli.StringFlag{
			Name:  "port, p",
			Value: "6667",
			Usage: "IRC port to connect to",
		},
		cli.StringFlag{
			Name:  "nickname, n",
			Value: "Jarvis",
			Usage: "Nickname to use",
		},
		cli.StringFlag{
			Name:  "channel",
			Value: "#jarvis-test",
			Usage: "Default IRC channel to join on connect",
		},
	}

	app.Action = start

	return app
}

func loadConfig(filename string) (*JarvisConfig, error) {

	config := new(JarvisConfig)

	if _, err := toml.DecodeFile(filename, config); err != nil {
		return nil, err
	}

	return config, nil
}

func newTLSConfig(server string, verify bool) *tls.Config {
	log.Debug("newTLSConfig for %s", server)
	config := new(tls.Config)
	config.ServerName = server
	config.InsecureSkipVerify = verify
	config.PreferServerCipherSuites = tlsPreferServerCipherSuites
	//config.CipherSuites = tlsCipherSuites

	return config
}
