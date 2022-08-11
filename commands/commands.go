package commands

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"

	"LogAnalyzer/config"
	"LogAnalyzer/logger"
)

type HandlerFunc func(in, out string, f *config.Filter) error

type command struct {
	name        string
	usage       string
	description string
	handlerFunc HandlerFunc
	mux         sync.Mutex
}

type Cmd interface {
	Name() string
	Usage() string
	Description() string
	write()
}

// New creates and registers a New command line command.
func New(name, usage, desc string) (c *command) {
	c = &command{
		name:        name,
		usage:       usage,
		description: desc,
	}
	return
}

func (c *command) Register(handler HandlerFunc) {
	c.mux.Lock()
	c.handlerFunc = handler
	c.mux.Unlock()

	reg.Lock()
	defer reg.Unlock()
	reg.commands = append(reg.commands, c)
	return
}

func (c *command) Execute(in, out string, f *config.Filter) (err error) {
	start := time.Now()
	err = c.handlerFunc(in, out, f)
	dur := time.Since(start)
	if err != nil {
		return
	}
	logger.Statsf("execution took %v", dur)
	return nil
}

// Name gets the name of this command.
func (c *command) Name() string {
	return c.name
}

// Usage gets the usage information of this command.
func (c *command) Usage() string {
	return c.usage
}

// Description gets the description of this command.
func (c *command) Description() string {
	return c.description
}

func (c *command) out(out string, d []byte) {
	if out == "" {
		fmt.Println(string(d))
		return
	}

	err := os.WriteFile(out, d, 0700)
	if err != nil {
		logger.Fatalf("unable to write file: %v\n", err)
	}
}

func (c *command) match(path string, f *config.Filter) (buf *bytes.Buffer, err error) {
	if f == nil {
		return nil, errors.New("filter cannot be nil")
	}
	_, err = os.Stat(path)
	if err != nil {
		return
	}

	cont, err := os.ReadFile(path)
	if err != nil {
		return
	}

	buf = &bytes.Buffer{}
	_, err = buf.Write(bytes.Join(regexp.MustCompile(f.Regex).FindAll(cont, -1), []byte("\n")))
	if buf.Len() < 1 {
		logger.Info("no matches found")
	}
	return
}
