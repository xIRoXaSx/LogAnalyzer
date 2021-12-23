package main

import (
	"LogAnalyzer/commands"
	"LogAnalyzer/configuration"
)

func main() {
	configuration.CreateConfigIfNotExists()
	configuration.ReadJson()
	commands.LoadCommands()
	CheckArgs()
}
