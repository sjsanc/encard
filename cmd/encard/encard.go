package main

import (
	"context"
	"log"
	"os"

	"github.com/sjsanc/encard/internal/encard"
	"github.com/urfave/cli/v3"
)

func run(args []string) error {
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
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "Print verbose output",
			},
		},
		Action: encard.DoRootAction,
	}

	return cmd.Run(context.Background(), args)
}

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
