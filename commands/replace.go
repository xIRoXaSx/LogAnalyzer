package commands

import (
	"errors"
	"os"
	"regexp"

	"LogAnalyzer/config"
)

const commandReplace = "replace"

type replace struct {
	*command
}

func init() {
	r := replace{
		New(
			commandReplace,
			"LogAnalyzer replace -filter <Filter> -in <Filepath> [-out <Filepath>]",
			"Replace all matched strings with the given replacement option",
		),
	}
	r.Register(r.Execute)
}

func (r *replace) Execute(in, out string, f *config.Filter) (err error) {
	if f == nil {
		return errors.New("filter cannot be nil")
	}
	_, err = os.Stat(in)
	if err != nil {
		return
	}

	cont, err := os.ReadFile(in)
	if err != nil {
		return
	}

	replaced := regexp.MustCompile(f.Regex).ReplaceAll(cont, []byte(f.Options.Replacement))
	r.out(out, replaced)
	return
}
