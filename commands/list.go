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

func init() {
	in := list{
		command: New(
			"list",
			"LogAnalyzer list",
			"List all configured filters",
		),
	}
	in.Register(in.Execute)
}

func (l *list) Execute(string, string, *config.Filter) (err error) {
	c := config.Get()
	var (
		maxLen    int
		filterLen int
	)
	for _, f := range c.Filters {
		le := len(f.Name)
		filterLen++
		if le > maxLen {
			maxLen = le
		}
	}

	iPadding := len(fmt.Sprint(filterLen))
	for i, f := range c.Filters {
		fLen := len(f.Name)
		logger.Infof("%*d: %s => %s\n", iPadding, i, f.Name+strings.Repeat(" ", maxLen-fLen), f.Regex)
	}
	return
}
