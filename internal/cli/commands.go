package cli

import (
	"context"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sjsanc/encard/internal/encard"
	"github.com/urfave/cli/v3"
)

func Run(args []string) error {
	cmd := &cli.Command{
		Name:      "encard",
		Usage:     "start the CLI",
		ArgsUsage: "[path/to/deck]",
		Action:    rootAction,
	}

	return cmd.Run(context.Background(), args)
}

var DEFAULT_ROOT_DECK = ".encard"

// TODO: match deck names directly with `encard <deckname>`
// Currently matches <path/to/file> | <path/to/dir> | HomeDir

// TODO: align the card answer beneath the card question

func rootAction(ctx context.Context, cmd *cli.Command) error {
	path := cmd.Args().First()

	if len(path) == 0 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("error getting home directory: %w", err)
		}
		path = homeDir + "/" + DEFAULT_ROOT_DECK
	}

	var cards []*encard.Card

	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("invalid path: %w", err)
	}

	if info.IsDir() {
		cards, err = encard.ParseCardsFromPath(path)
		if err != nil {
			return fmt.Errorf("error parsing cards: %w", err)
		}
	}

	if strings.HasSuffix(path, ".md") {
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}
		cards = encard.ParseCards(string(data), path)
	}

	model := &encard.Model{
		Cards:        cards,
		CurrentIndex: 0,
	}

	program := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		return err
	}
	return nil
}
