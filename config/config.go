package config

import (
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/xIRoXaSx/gonfig"
)

const (
	Dir  = "LogAnalyzer"
	File = "config"
)

type Config struct {
	PrintStats bool     `yaml:"PrintStats"`
	Filters    []Filter `yaml:"Filters"`
}

type Filter struct {
	Name    string `yaml:"Name"`
	Regex   string `yaml:"Regex"`
	Options *Opts  `yaml:"Options,omitempty"`
}

type Opts struct {
	Replacement      string `yaml:"Replacement"`
	RemoveEmptyLines bool   `yaml:"RemoveEmptyLines"`
}

func Default() *Config {
	return &Config{
		PrintStats: true,
		Filters: []Filter{
			{
				Name:  "JsonMin",
				Regex: "(\\s+[^{}\"'\\[\\]\\\\\\w])|(\\B\\s)",
				Options: &Opts{
					Replacement:      "",
					RemoveEmptyLines: true,
				},
			},
			{
				Name:  "JavaStackTrace",
				Regex: "(?m)^.*?Exception.*(?:[\\r|\\n]+^\\s*at .*)+",
			},
			{
				Name:  "Exception",
				Regex: "(?m)((.*(\\n|\\r|\\r\\n)){1})^.*?Exception.*(?:[\\n|\\r|\\r\\n]+^\\s*at .*)+",
				Options: &Opts{
					Replacement:      "Nope, not a single error to report ;)",
					RemoveEmptyLines: true,
				},
			},
		},
	}
}

func Get() (c *Config) {
	dir := Dir
	if runtime.GOOS != "windows" {
		dir = strings.ToLower(Dir)
	}

	g, err := gonfig.New(dir, File, gonfig.GonfYAML, false)
	if err != nil {
		log.Fatalf("unable to create config: %v\n", err)
	}

	_, err = os.Stat(g.FullPath())
	if os.IsNotExist(err) {
		err = g.WriteToFile(Default())
		if err != nil {
			log.Fatalf("unable to write config: %v\n", err)
		}
		log.Printf("Configuration has been written to %s\n", g.FullPath())
		os.Exit(0)
	}

	err = g.Load(&c)
	if err != nil {
		log.Fatalf("unable to read config file: %v\n", err)
	}
	return
}

// RetrieveFilter gets the defined filter from the config.
func (c *Config) RetrieveFilter(name string) (f *Filter) {
	for _, filter := range c.Filters {
		if filter.Name == name {
			f = &filter
			break
		}
	}
	return
}
