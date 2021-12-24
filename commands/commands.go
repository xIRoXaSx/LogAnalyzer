package commands

import (
	"LogAnalyzer/helper"
	"LogAnalyzer/logger"
	"LogAnalyzer/structs"
	"fmt"
	survey "github.com/AlecAivazis/survey/v2"
	"os"
	"path/filepath"
	"strings"
)

var Commands []structs.Command

func putCommand(name string, usage string, description string) {
	Commands = append(Commands, structs.Command{Name: name, Usage: usage, Description: description})
}

// LoadCommands stores commands in the Commands variable
func LoadCommands() {
	putCommand("help", "LogAnalyzer h[elp]", "Show help for each command")
	putCommand("inspect", "LogAnalyzer i[nspect] <Filter name> <Logfile>", "Inspect a logfile with the provided filter")
	putCommand("listfilter", "LogAnalyzer l[istfilter]", "List all configured filters")
	putCommand("replace", "LogAnalyzer r[eplace] <Filter name> <Logfile> [Replacement]",
		"Replace all matched occurrences from the provided filter with [Replacement] or the corresponding value from the config")
}

// PrintCommands prints all usable commands
func PrintCommands() {
	textAvailCmd := "Commands"
	textUsage := "Usage"
	textDescription := "Description"
	maxLengthName, maxLengthUsage, maxLengthDescription := helper.GetLongestStringCommand(Commands)

	fmt.Println(
		textAvailCmd + helper.GetSpaceSeparator(len(textAvailCmd), maxLengthName, " ") +
			textUsage + helper.GetSpaceSeparator(len(textUsage), maxLengthUsage, " ") +
			textDescription + helper.GetSpaceSeparator(len(textDescription), maxLengthDescription, " "))

	fmt.Println(
		helper.GetSpaceSeparator(1, maxLengthName, "-") +
			helper.GetSpaceSeparator(1, maxLengthUsage, "-") +
			helper.GetSpaceSeparator(1, maxLengthDescription, "-"))

	for i := 0; i < len(Commands); i++ {
		fmt.Println(
			Commands[i].Name + helper.GetSpaceSeparator(maxLengthName, len(Commands[i].Name), " ") +
				Commands[i].Usage + helper.GetSpaceSeparator(maxLengthUsage, len(Commands[i].Usage), " ") +
				Commands[i].Description + helper.GetSpaceSeparator(maxLengthDescription, len(Commands[i].Description), " "))
	}
}

// CompleteFilterPrompt prompts the user for a filter to use
func CompleteFilterPrompt(filters []structs.Filter) (string, error) {
	var filterName []string
	var response string
	for i := 0; i < len(filters); i++ {
		filterName = append(filterName, filters[i].Name)
	}

	prompt := &survey.Select{
		Message: "Choose a filter: ",
		Help:    "The filter to use for the defined action",
		Options: filterName,
	}

	err := survey.AskOne(prompt, &response)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return response, nil
}

// CompleteCommandPrompt prompts the user for a command to use
func CompleteCommandPrompt() (string, error) {
	var commandName []string
	var response string
	for i := 0; i < len(Commands); i++ {
		commandName = append(commandName, Commands[i].Name)
	}

	prompt := &survey.Select{
		Message: "Choose a command: ",
		Help:    "The command to execute",
		Options: commandName,
	}

	err := survey.AskOne(prompt, &response)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return response, nil
}

// CompleteFilePrompt prompts the user for a file to use
func CompleteFilePrompt() (string, error) {
	var response string
	prompt := &survey.Input{
		Message: "Choose a file: ",
		Help:    "The file to manipulate. ",
		Suggest: getFilesAndDirs,
	}

	err := survey.AskOne(prompt, &response)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return response, nil
}

// getFilesAndDirs gets a slice of directory and file names.
// Directories of the returned slice will have a "/" appended to it
func getFilesAndDirs(toComplete string) []string {
	files, _ := filepath.Glob(toComplete + "*")
	for i := 0; i < len(files); i++ {
		file, err := os.Lstat(files[i])
		if err != nil {
			continue
		}

		if file.IsDir() {
			files[i] = toComplete[:strings.LastIndex(toComplete, string(os.PathSeparator))] +
				string(os.PathSeparator) + file.Name() + string(os.PathSeparator)
			continue
		}
	}

	return files
}
