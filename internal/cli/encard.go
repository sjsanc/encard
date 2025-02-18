package cli

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sjsanc/encard/internal/cards"
	"github.com/sjsanc/encard/internal/encard"
	"github.com/urfave/cli/v3"
)

func encardAction(ctx context.Context, cmd *cli.Command) error {
	cfg, err := encard.NewConfig(cmd.String("config"))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	args := cmd.Args().Slice()

	matches := []string{}

	filepath.WalkDir(cfg.CardsDir, func(path string, d fs.DirEntry, err error) error {
		for _, arg := range args {
			// TODO: handle filepath.Match error
			matched, _ := filepath.Match(arg, d.Name())
			if matched {
				matches = append(matches, path)
			}
		}
		return nil
	})

	var cards []cards.Card

	for _, match := range matches {
		deck, err := encard.LoadDeckFromPath(match)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		cards = append(cards, deck...)
	}

	model := &encard.Model{
		Cards: cards,
	}

	if cmd.Bool("shuffle") {
		model.IsShuffled = true
		// model.Cards = encard.Shuffle(model.Cards)
	}

	program := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		return err
	}
	return nil
}
