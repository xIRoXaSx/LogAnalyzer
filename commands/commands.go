package commands

import (
	"LogAnalyzer/helper"
	"LogAnalyzer/structs"
	"fmt"
)

var Commands []structs.Command

func putCommand(name string, usage string, description string) {
	Commands = append(Commands, structs.Command{Name: name, Usage: usage, Description: description})
}

func LoadCommands() {
	putCommand("listfilter", "LogAnalyzer l[istfilter]", "List all configured filters")
	putCommand("inspect", "LogAnalyzer i[nspect] <Filter name> <Logfile>", "Inspect a logfile with the provided filter")
}

func PrintCommands() {
	if len(Commands) < 1 {
		LoadCommands()
	}

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
