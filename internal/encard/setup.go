package encard

import (
	"fmt"

	"github.com/urfave/cli/v3"
)

var logger = NewLogger(false)

// `Options` contains runtime options.`
type Options struct {
	shuffled bool
	verbose  bool
	cfg      *Config
}

func Setup(cmd *cli.Command) (*Options, error) {
	cfgPath := cmd.String("config")

	cfg, err := NewConfig(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	opts := &Options{
		cfg:      cfg,
		shuffled: cmd.Bool("shuffle"),
		verbose:  cmd.Bool("verbose"),
	}

	if opts.verbose {
		logger = NewLogger(true)
	}

	if cfgPath != "" {
		logger.Printf("using configuration from %s", cfgPath)
	} else {
		logger.Printf("using default configuration")
	}

	logger.Printf("using %s as default load path", cfg.CardsDir)

	return opts, nil
}
