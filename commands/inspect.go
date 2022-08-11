package commands

import (
	"LogAnalyzer/config"
)

type inspect struct {
	*command
}

func init() {
	in := inspect{
		command: New(
			"inspect",
			"LogAnalyzer inspect -filter <Filter> -in <Filepath> [-out <Filepath>]",
			"Inspect a logfile with the provided filter",
		),
	}
	in.Register(in.Execute)
}

func (ins *inspect) Execute(in, out string, f *config.Filter) (err error) {
	buf, err := ins.match(in, f)
	if err != nil {
		return
	}

	ins.out(out, buf.Bytes())
	return
}
