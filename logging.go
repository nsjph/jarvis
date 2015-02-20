package main

import (
	"github.com/op/go-logging"
	"os"
)

var (
	log              = logging.MustGetLogger("jarvis")
	logBackendStderr = logging.NewLogBackend(os.Stderr, "", 0)
	logFormat        = logging.MustStringFormatter(
		"%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}")
)

func initLogging() {
	logging.SetFormatter(logFormat)
	logging.SetBackend(logBackendStderr)
	log.Debug("Logging configured")
}
