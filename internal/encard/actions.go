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

	cards, err := LoadCards(cfg.CardsDir, globs)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if len(cards) == 0 {
		return fmt.Errorf("no cards found")
	}

	opts := &Opts{
		shuffled: cmd.Bool("shuffle"),
		verbose:  cmd.Bool("verbose"),
	}

	session := NewSession(cards, opts)

	model := NewModel(session)

	program := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		return err
	}
	return nil
}
