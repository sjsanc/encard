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

func rootAction(context.Context, *cli.Command) error {
	path := "decks"
	if len(path) == 0 {
		return fmt.Errorf("no path provided")
	}

	var cards []*encard.Card
	var err error

	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("invalid path: %w", err)
	}

	if info.IsDir() {
		cards, err = encard.ParseCardsFromPath(path)
	} else if strings.HasSuffix(path, ".md") {
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}
		cards = encard.ParseCards(string(data), path)
	} else {
		return fmt.Errorf("invalid path")
	}

	if err != nil {
		return err
	}

	model := &encard.Model{
		Cards:        cards,
		CurrentIndex: 0,
	}

	fmt.Println("Loaded", len(cards), "cards.")

	program := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		return err
	}
	return nil
}
