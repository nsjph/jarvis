package main

import (
	glispext "github.com/zhemao/glisp/extensions"
	"github.com/zhemao/glisp/interpreter"
	"os"
)

// These are the equivalent of the plugin library. These are the
// variables and functions that are assumed present in every plugin,
// so plugin authors dont have to keep reinventing the wheel
// These are loaded after regexp extension is loaded
var glispPluginFunctions string = `
(def command false)
(def pattern false)
(def prefix false)
(def regexp false)
(defn update-regexp [] 
  (def regexp 
    (regex-compile (apply "^" (apply prefix pattern)))))
(defn update-prefix [p]
  (def prefix p))
`

type GlispPlugin struct {
	prefix   string
	command  string
	pattern  string
	filename string
	env      *glisp.Glisp
}

func newGlispEnv() *glisp.Glisp {

	env := glisp.NewGlisp()
	glispext.ImportRegex(env)
	err := env.LoadString(glispPluginFunctions)
	if err != nil {
		log.Fatalf("Error loading glisp functions: %v", err)
	}

	return env
}

func getCommand(e *glisp.Glisp) string {
	sexp, success := e.FindObject("command")
	if success != true {
		log.Error("Failed to find object \"command\"")
		return ""
	} else {
		log.Debug(sexp.SexpString())
	}
	return sexp.SexpString()

}

func (c *Connection) registerGlispPluginHandler(p *GlispPlugin) (err error) {

	command, err := p.env.EvalString("command")
	if err != nil {
		log.Debug("Unable to obtain command from plugin")
		return err
	} else {
		p.command = command.SexpString()
	}
	return nil

}

func loadGlispPlugins(pluginFiles []string) ([]*GlispPlugin, error) {

	var plugins []*GlispPlugin

	for _, v := range pluginFiles {
		log.Debug("Attempting to initialize %s", v)

		file, err := os.Open(v)
		if err != nil {
			log.Error("Failed to open %s: %v", v, err)
			break
		}
		defer file.Close()

		p := new(GlispPlugin)
		p.env = newGlispEnv()

		err = p.env.LoadFile(file)
		if err != nil {
			log.Error("Failed to load %s: %v", v, err)
		} else {
			p.prefix = "@"
			plugins = append(plugins, p)
			log.Debug("plugin command is %s", getCommand(p.env))
			ev, _ := p.env.EvalString("command")
			log.Debug("eval string %s", ev.SexpString())
			log.Info("Loaded plugin: %s", v)
		}

	}

	return plugins, nil
}
