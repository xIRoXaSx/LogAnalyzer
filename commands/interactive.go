package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"LogAnalyzer/config"
	"LogAnalyzer/logger"
	sv "github.com/AlecAivazis/survey/v2"
)

// Interactive keeps the user inside an interactive terminal environment until the user decides to close the program.
func Interactive() {
	const exitOpt = "exit"
	var (
		maxNameLen int
		cmdLen     int
	)
	commands := reg.commands
	for _, c := range commands {
		nameLen := len(c.name)
		cmdLen++
		if nameLen > maxNameLen {
			maxNameLen = nameLen
		}
	}

	// Format commands for the selection.
	opts := make([]string, 0)
	for _, c := range commands {
		name := c.Name()
		opts = append(opts, fmt.Sprintf("%s | %s", name+strings.Repeat(" ", maxNameLen-len(name)), c.Description()))
	}
	opts = append(opts, fmt.Sprintf("%s | %s", exitOpt+strings.Repeat(" ", maxNameLen-len(exitOpt)), "Exit LogAnalyzer"))

	var resp int
	for {
		err := sv.AskOne(&sv.Select{
			Message: "Chose a command: ",
			Options: opts,
			Help:    "Select a command to execute with the [ENTER] key.",
		}, &resp)
		if err != nil {
			logger.Fatal(err)
		}
		if resp == len(opts)-1 {
			return
		}

		cmd := commands[resp]
		if err != nil {
			logger.Error(err)
			continue
		}

		switch cmd.Name() {
		case commandReplace, commandInspect:
			fs := config.Get().Filters
			var ind int
			ind, err = askForFilter(fs)
			if err != nil {
				logger.Error(err)
				continue
			}

			var (
				in  string
				out string
			)
			in, out, err = askForFilePaths()
			if err != nil {
				logger.Errorf("\n%v\n", err)
				continue
			}
			err = cmd.Execute(in, out, &fs[ind])
			if err != nil {
				logger.Errorf("\n%v\n", err)
			}
		default:
			err = cmd.Execute("", "", nil)
			if err != nil {
				logger.Error(err)
			}
		}
	}
}

func askForFilter(fs []config.Filter) (ind int, err error) {
	filters := formatFilters(fs)
	err = sv.AskOne(&sv.Select{
		Message: "Chose the filter: ",
		Options: filters,
		Help:    "Select one of the defined filters from your configuration file.",
	}, &ind)
	return
}

func askForFilePaths() (in, out string, err error) {
	ask := func(msgOpt string) (resp string, err error) {
		err = sv.AskOne(&sv.Input{
			Message: fmt.Sprintf("Chose the %s file: ", msgOpt),
			Suggest: getPath,
		}, &resp)
		return
	}
	in, err = ask("input")
	if err != nil {
		return
	}
	out, err = ask("output")
	return
}

// getPath gets a slice of directory and file names.
// Directories will end with os.PathSeparator.
func getPath(p string) []string {
	paths, _ := filepath.Glob(p + "*")
	pLength := len(paths)
	for i := 0; i < pLength; i++ {
		f, err := os.Lstat(paths[i])
		if err != nil {
			continue
		}

		// Autocomplete directory names.
		if f.IsDir() {
			lastInd := strings.LastIndex(p, string(os.PathSeparator))
			if lastInd == -1 {
				paths[i] = p + f.Name()[len(p):] + string(os.PathSeparator)
				continue
			}
			paths[i] += string(os.PathSeparator)
		}
	}
	return paths
}
