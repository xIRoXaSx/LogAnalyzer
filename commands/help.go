package commands

import (
	"fmt"
	"strings"

	"LogAnalyzer/config"
	"LogAnalyzer/logger"
)

type help struct {
	*command
}

func init() {
	in := help{
		command: New(
			"help",
			"LogAnalyzer help",
			"Show this help message",
		),
	}
	in.Register(in.Execute)
}

func (h *help) Execute(string, string, *config.Filter) (err error) {
	var (
		maxNameLen  int
		maxUsageLen int
		cmdLen      int
	)
	cmds := reg.commands
	for _, c := range cmds {
		nameLen := len(c.name)
		usageLen := len(c.usage)
		cmdLen++
		if nameLen > maxNameLen {
			maxNameLen = nameLen
		}
		if usageLen > maxUsageLen {
			maxUsageLen = usageLen
		}
	}

	iPadding := len(fmt.Sprint(cmdLen))
	for i, c := range cmds {
		u := c.Usage()
		n := c.Name()
		uLen := len(u)
		nLen := len(n)
		logger.Infof(
			"%*d: %s | %s | %s\n",
			iPadding, i, n+strings.Repeat(" ", maxNameLen-nLen),
			u+strings.Repeat(" ", maxUsageLen-uLen),
			c.Description(),
		)
	}

	return
}
