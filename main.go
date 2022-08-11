package main

import (
	"flag"
	"os"

	"LogAnalyzer/commands"
	"LogAnalyzer/config"
	"LogAnalyzer/logger"
)

var conf *config.Config

func init() {
	conf = config.Get()
	logger.New(conf.PrintStats)
}

func main() {
	var (
		in      string
		out     string
		noStats bool
	)

	args := os.Args
	if len(args) == 1 {
		logger.Fatal("no command supplied")
	}

	cmd := args[1]
	c, err := commands.Retrieve(cmd)
	if err != nil {
		c, err = commands.Retrieve("help")
		if err != nil {
			os.Exit(0)
		}
		_ = c.Execute("", "", nil)
		return
	}

	var filter string
	os.Args = args[2:]
	flag.StringVar(&in, "in", "", "the file to analyze")
	flag.StringVar(&out, "out", "", "the output file. If no file path given, stdout will be used instead")
	flag.StringVar(&filter, "filter", "", "the filter to use")
	flag.BoolVar(&noStats, "nostats", false, "disable stats message at the end of the execution")
	err = flag.CommandLine.Parse(os.Args)
	if err != nil {
		logger.Fatalf("unable to parse args: %v\n", err)
	}

	f := conf.RetrieveFilter(filter)
	err = c.Execute(in, out, f)
	if err != nil {
		logger.Errorf("handler error: %v\n", err)
	}
}
