package commands

import (
	"fmt"
	"strings"

	"LogAnalyzer/config"
	"LogAnalyzer/logger"
)

type list struct {
	*command
}

const commandList = "list"

func init() {
	in := list{
		command: New(
			commandList,
			"LogAnalyzer list",
			"List all configured filters",
		),
	}
	in.Register(in.Execute)
}

func (l *list) Execute(string, string, *config.Filter) (err error) {
	filters := formatFilters(config.Get().Filters)
	for _, f := range filters {
		logger.Info(f)
	}
	return
}

func formatFilters(fs []config.Filter) (filters []string) {
	var (
		maxLen    int
		filterLen int
	)
	for _, f := range fs {
		le := len(f.Name)
		filterLen++
		if le > maxLen {
			maxLen = le
		}
	}
	iPadding := len(fmt.Sprint(maxLen))
	filters = make([]string, len(fs))
	for i, f := range fs {
		fLen := len(f.Name)
		filters[i] = fmt.Sprintf("%*d: %s => %s", iPadding, i, f.Name+strings.Repeat(" ", maxLen-fLen), f.Regex)
	}
	return
}
