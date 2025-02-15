package cli

import (
	"context"

	"github.com/urfave/cli/v3"
)

func Run(args []string) error {
	cmd := &cli.Command{
		Name:      "encard",
		Usage:     "start the CLI",
		ArgsUsage: "[path/to/deck]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "shuffle",
				Aliases: []string{"s"},
				Usage:   "Shuffle the cards before starting",
			},
		},
		Action: encardAction,
	}

	return cmd.Run(context.Background(), args)
}
