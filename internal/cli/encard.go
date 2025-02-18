package cli

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gobwas/glob"
	"github.com/sjsanc/encard/internal/encard"
	"github.com/urfave/cli/v3"
)

func encardAction(ctx context.Context, cmd *cli.Command) error {
	cfg, err := encard.NewConfig(cmd.String("config"))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	args := cmd.Args().Slice()

	globs := make([]glob.Glob, len(args))
	for i, arg := range args {
		globs[i] = glob.MustCompile(arg)
	}

	cards, err := encard.LoadCards(cfg.CardsDir, globs)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if len(cards) == 0 {
		return fmt.Errorf("no cards found")
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
