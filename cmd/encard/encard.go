package main

import (
	"context"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sjsanc/encard/internal/encard"
	"github.com/sjsanc/encard/internal/tui"
	"github.com/urfave/cli/v3"
)

func run(args []string) error {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:      "run",
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
				Action: func(ctx context.Context, c *cli.Command) error {
					opts, err := encard.Setup(c)
					if err != nil {
						log.Fatalf("%v", err)
					}

					args := c.Args().Slice()

					cards, _ := encard.LoadCards(args, opts.Cfg.CardsDir)

					if len(cards) == 0 {
						return fmt.Errorf("no cards found")
					}

					session := encard.NewSession(cards, opts)
					model := tui.NewModel(session)

					program := tea.NewProgram(model, tea.WithAltScreen())
					if _, err := program.Run(); err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:      "verify",
				Usage:     "verify the parsed cards",
				ArgsUsage: "[path/to/deck]",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "verbose",
						Aliases: []string{"v"},
						Usage:   "Print verbose output",
					},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					opts, err := encard.Setup(c)
					if err != nil {
						log.Fatalf("%v", err)
					}
					args := c.Args().Slice()
					cards, _ := encard.LoadCards(args, opts.Cfg.CardsDir)

					if len(cards) == 0 {
						return fmt.Errorf("no cards found")
					}

					fmt.Printf("loaded %d cards\n", len(cards))

					return nil
				},
			},
		},
	}

	return cmd.Run(context.Background(), args)
}

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
