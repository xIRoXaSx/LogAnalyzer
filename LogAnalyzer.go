package main

import "LogAnalyzer/configuration"

func main() {
	configuration.CreateConfigIfNotExists()
	configuration.ReadJson()
}
