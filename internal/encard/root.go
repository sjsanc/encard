package encard

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gobwas/glob"
	"github.com/urfave/cli/v3"
)

func DoRootAction(ctx context.Context, cmd *cli.Command) error {
	cfg, err := NewConfig(cmd.String("config"))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	args := cmd.Args().Slice()

	globs := make([]glob.Glob, len(args))
	for i, arg := range args {
		globs[i] = glob.MustCompile(arg)
	}

	decks, err := LoadCards(cfg.CardsDir, globs)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if len(decks) == 0 {
		return fmt.Errorf("no cards found")
	}

	keys := make([]string, 0, len(decks))
	for k := range decks {
		keys = append(keys, k)
	}
	first := keys[0]

	model := &Model{
		DeckMap:        decks,
		CurrentDeckKey: first,
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
