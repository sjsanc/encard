package encard

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v3"
)

func Run(ctx context.Context, cmd *cli.Command) error {
	opts := Setup(cmd)
	args := cmd.Args().Slice()
	cards, _ := LoadCards(args, opts.cfg)

	if len(cards) == 0 {
		return fmt.Errorf("no cards found")
	}

	session := NewSession(cards, opts)
	model := NewModel(session)

	program := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		return err
	}
	return nil
}

func Verify(ctx context.Context, cmd *cli.Command) error {
	opts := Setup(cmd)
	args := cmd.Args().Slice()
	cards, _ := LoadCards(args, opts.cfg)

	if len(cards) == 0 {
		return fmt.Errorf("no cards found")
	}

	fmt.Printf("loaded %d cards\n", len(cards))

	return nil
}
