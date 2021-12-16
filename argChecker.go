package main

import (
	"LogAnalyzer/commands"
	"LogAnalyzer/configuration"
	"LogAnalyzer/helper"
	"LogAnalyzer/logger"
	"LogAnalyzer/structs"
	"fmt"
	"os"
	"regexp"
)

// CheckArgs checks the os.Args for a passed Filter
func CheckArgs() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	inspectRegex := regexp.MustCompile(`^i[nspect]?`)
	listFilterRegex := regexp.MustCompile(`^f[ilters]?|^l[istfler]?`)
	replaceFilterRegex := regexp.MustCompile(`^r[eplac]?`)

	switch {
	case inspectRegex.MatchString(os.Args[1]):
		filter, filePath := getFilterAndFilePathFromArgs(os.Args[1:])

		logger.Info("Used filter \"" + filter.Name + "\"")
		commands.Inspect(filePath, filter)
		return
	case listFilterRegex.MatchString(os.Args[1]):
		commands.ListFilter()
		return
	case replaceFilterRegex.MatchString(os.Args[1]):
		filter, filePath := getFilterAndFilePathFromArgs(os.Args[1:])
		replacement := filter.Replacement

		if len(os.Args) >= 5 {
			replacement = os.Args[4]
		}

		commands.Replace(filePath, filter, replacement)
		return
	default:
		printHelp()
	}

	return
}

// getFilterAndFilePathFromArgs checks each argument for filter name and file path.
// This enables users to swap filter names and file paths.
func getFilterAndFilePathFromArgs(args []string) (structs.Filter, string) {
	filter := structs.Filter{}
	filePath := ""

	for i := 0; i < len(args); i++ {
		f := helper.ContainsFilterName(configuration.JsonConfig.LogAnalyzer.Filters, args[i])
		if filter == (structs.Filter{}) && f != (structs.Filter{}) {
			filter = f
			continue
		}

		if _, err := os.Stat(args[i]); err == nil && filePath == "" {
			filePath = args[i]
			continue
		}
	}

	return filter, filePath
}

// printHelp prints all available commands, their usage and description
func printHelp() {
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  LogAnalyzer [command] [arguments]")
	fmt.Println("")
	commands.PrintCommands()
	fmt.Println("")
}
