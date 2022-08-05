package main

import (
	"log"
	"os"

	"LogAnalyzer/logger"

	"github.com/xIRoXaSx/gonfig"
)

func init() {
	g, err := gonfig.New(ConfigDir, ConfigFile, gonfig.GonfYAML, false)
	if err != nil {
		log.Fatalf("unable to create config: %v\n", err)
	}

	_, err = os.Stat(g.FullPath())
	if os.IsNotExist(err) {
		err = g.WriteToFile(defaultConfig())
		if err != nil {
			log.Fatalf("unable to write config: %v\n", err)
		}
		log.Printf("Configuration has been written to %s\n", g.FullPath())
		os.Exit(0)
	}

	err = g.Load(conf)
	if err != nil {
		log.Fatalf("unable to read config file: %v\n", err)
	}

	logger.New(conf.Debug)
}

func main() {

}
