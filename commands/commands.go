package commands

import "sync"

type registry struct {
	commands []Command
	sync.Mutex
}

var reg registry

type Command struct {
	name        string
	usage       string
	description string
	Cmd
}

type Cmd interface {
	Name() string
	Usage() string
	Description() string
	Execute() error
}

// New creates and registers a new command line command.
func New(name, usage, desc string) {
	c := Command{
		name:        name,
		usage:       usage,
		description: desc,
	}
	reg.Lock()
	defer reg.Unlock()
	reg.commands = append(reg.commands, c)
	return
}

// Name gets the name of this command.
func (c *Command) Name() string {
	return c.name
}

// Usage gets the usage information of this command.
func (c *Command) Usage() string {
	return c.usage
}

// Description gets the description of this command.
func (c *Command) Description() string {
	return c.description
}
