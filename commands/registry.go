package commands

import (
	"errors"
	"strings"
	"sync"
)

var ErrNotExist = errors.New("command does not exist")

type registry struct {
	commands []*command
	sync.Mutex
}

var reg registry

func Retrieve(r string) (c *command, err error) {
	lower := strings.ToLower(r)
	for i := 0; i < len(reg.commands); i++ {
		c = reg.commands[i]
		if lower != c.name {
			continue
		}
		return c, nil
	}
	return nil, ErrNotExist
}
