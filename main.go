package main

import (
	"os"
	"time"
)

func main() {

	initLogging()

	app := newApp()
	go app.Run(os.Args)

	sleepInterval := 60
	for {
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}

}
