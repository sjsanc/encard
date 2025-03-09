package encard

import "github.com/urfave/cli/v3"

var logger = NewLogger(false)

// `Options` contains runtime options.`
type Options struct {
	shuffled bool
	verbose  bool
	cfg      *Config
}

func Setup(cmd *cli.Command) *Options {
	cfg, err := NewConfig(cmd.String("config"))
	if err != nil {

	}

	opts := &Options{
		cfg:      cfg,
		shuffled: cmd.Bool("shuffle"),
		verbose:  cmd.Bool("verbose"),
	}

	if opts.verbose {
		logger = NewLogger(true)
	}

	return opts
}
