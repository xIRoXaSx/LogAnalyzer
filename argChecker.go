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
	// If no arguments are provided use interactive mode
	if len(os.Args) == 1 {
		selectedCommand, err := commands.CompleteCommandPrompt()
		if err == nil {
			os.Args = append(os.Args, selectedCommand)
		} else {
			logger.Critical(err.Error())
			return
		}

		if selectedCommand == "help" {
			printHelp()
			return
		}

		selectedFilter, err := commands.CompleteFilterPrompt(configuration.JsonConfig.LogAnalyzer.Filters)
		if err == nil {
			os.Args = append(os.Args, selectedFilter)
		} else {
			logger.Critical(err.Error())
			return
		}

		selectedFile, err := commands.CompleteFilePrompt()
		if err == nil {
			os.Args = append(os.Args, selectedFile)
		} else {
			logger.Critical(err.Error())
			return
		}
	}

	helpRegex := regexp.MustCompile(`^h[elp]?`)
	inspectRegex := regexp.MustCompile(`^i[nspect]?`)
	listFilterRegex := regexp.MustCompile(`^f[ilters]?|^l[istfler]?`)
	replaceFilterRegex := regexp.MustCompile(`^r[eplac]?`)
	filter, filePath := getFilterAndFilePathFromArgs(os.Args[1:])

	switch {
	case len(os.Args) == 1 || helpRegex.MatchString(os.Args[1]):
		printHelp()
		return
	case inspectRegex.MatchString(os.Args[1]):
		logger.Info("Used filter \"" + filter.Name + "\"")
		commands.Inspect(filePath, filter)
		return
	case listFilterRegex.MatchString(os.Args[1]):
		commands.ListFilter()
		return
	case replaceFilterRegex.MatchString(os.Args[1]):
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
