package encard

import (
	"fmt"

	"github.com/sjsanc/encard/internal/log"
	"github.com/urfave/cli/v3"
)

// `Options` contains runtime options.`
type Options struct {
	Cfg      *Config
	Shuffled bool
	Verbose  bool
}

func Setup(cmd *cli.Command) (*Options, error) {
	cfgPath := cmd.String("config")

	cfg, err := NewConfig(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	opts := &Options{
		Cfg:      cfg,
		Shuffled: cmd.Bool("shuffle"),
		Verbose:  cmd.Bool("verbose"),
	}

	if opts.Verbose {
		log.VERBOSE = true
	}

	if cfgPath != "" {
		log.Info("using configuration from %s", cfgPath)
	} else {
		log.Info("using default configuration")
	}

	log.Info("using %s as default load path", cfg.CardsDir)

	return opts, nil
}
