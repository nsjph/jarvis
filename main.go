package main

import (
	_ "fmt"
	cli "github.com/jawher/mow.cli"
	_ "github.com/pkg/profile"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var ()

func createConfigDirectory(configPath string) {
	err := os.Mkdir(configPath, 0700)
	if err != nil {
		panic(err)
	}
}

func initConfigDirectory(configDir string) {

	//configPath := filepath.Join(os.Getenv("HOME"), ".jarvis")

	_, err := os.Stat(configDir)
	if err != nil {
		createConfigDirectory(configDir)
	}

	// // check if the source is indeed a directory or not
	// if !src.IsDir() {
	// 	fmt.Printf("Error: %s is not a directory!\n", configDir)
	// 	os.Exit(1)
	// }
}

func main() {

	initLogging()

	app := cli.App("jarvis", "A helpful IRC bot")
	configDir := app.StringOpt("d dir", filepath.Join(os.Getenv("HOME"), ".jarvis"), "Configuration directory")

	app.Action = func() {

		initConfigDirectory(*configDir)

		configPath := filepath.Join(*configDir, "jarvis.toml")

		config, err := loadConfig(configPath)
		if err != nil {
			log.Fatalf("Unable to load config file %s: %v", configPath, err)
		} else {
			config.filename = configPath
		}

		jarvis := &Jarvis{
			connections: make(map[string]*Connection),
			config:      config,
			db:          initDatabase(*configDir),
		}

		jarvis.initConnections()
		for k, v := range jarvis.connections {
			err := v.startConnection()
			if err != nil {
				log.Error("Error connecting to %s: %v", k, err)
			}
		}
	}

	app.Run(os.Args)

	sleepInterval := 60
	for {
		memstats := new(runtime.MemStats)
		runtime.ReadMemStats(memstats)
		log.Debug("Memory allocated: %d", memstats.Alloc)
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}

}
