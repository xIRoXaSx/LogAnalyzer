package main

import (
	"LogAnalyzer/configuration"
	"LogAnalyzer/helper"
	"fmt"
	"os"
	"strings"
)

// CheckArgs checks the os.Args for a passed Filter
func CheckArgs() configuration.Filter {
	var filter configuration.Filter
	if len(os.Args) < 2 {
		return configuration.Filter{}
	}

	for i := 0; i < len(os.Args); i++ {
		if filter == (configuration.Filter{}) {
			filter = helper.Contains(configuration.JsonConfig.LogAnalyzer.Filters, os.Args[i])
		}

		switch strings.ToLower(os.Args[i]) {
		case "listfilter":
		default:
			printHelp()
		}
	}

	return filter
}

// printHelp prints all available commands, their usage and description
func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  LogAnalyzer [command]")
	fmt.Println("")
	fmt.Println("Available commands\tUsage\t\t\t\t\t\tDescription")
	fmt.Println("  inspect\t\tLogAnalyzer i[nspect] <Filter name> <Logfile>\tInspect a logfile with the provided filter")
}
