package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sjsanc/encard/internal/encard"
	"github.com/urfave/cli/v3"
)

var ENCARD_FOLDER_ROOT = ".encard"

func encardAction(ctx context.Context, cmd *cli.Command) error {
	args := cmd.Args().Slice()

	// Determine the base directory for card files
	cardDirPath := os.Getenv("XDG_DATA_HOME")
	if cardDirPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("home directory is not set: %w", err)
		}
		cardDirPath = filepath.Join(homeDir, ENCARD_FOLDER_ROOT)
	} else {
		cardDirPath = filepath.Join(cardDirPath, ENCARD_FOLDER_ROOT)
	}

	cardDir, err := os.Stat(cardDirPath)
	if err != nil || !cardDir.IsDir() {
		return fmt.Errorf("invalid card directory path: %w", err)
	}

	var cards []*encard.Card

	if len(args) == 0 {
		cards, err = encard.ParseCardsFromPath(cardDirPath)
		if err != nil {
			return fmt.Errorf("error parsing cards: %w", err)
		}
		return nil
	}

	for _, arg := range args {
		var fullPath string

		if filepath.IsAbs(arg) {
			fullPath = arg
		} else {
			fullPath = filepath.Join(cardDirPath, arg)
		}

		info, err := os.Stat(fullPath)
		if err != nil {
			return fmt.Errorf("invalid path: %w", err)
		}

		if info.IsDir() {
			parsed, err := encard.ParseCardsFromPath(fullPath)
			if err != nil {
				return fmt.Errorf("error parsing cards from directory %s: %w", fullPath, err)
			}
			cards = append(cards, parsed...)
			continue
		}

		if filepath.Ext(fullPath) == ".md" {
			data, err := os.ReadFile(fullPath)
			if err != nil {
				return fmt.Errorf("error reading file %s: %w", fullPath, err)
			}
			parsed := encard.ParseCards(string(data), fullPath)
			cards = append(cards, parsed...)
		}
	}

	model := &encard.Model{
		Cards: cards,
	}

	if cmd.Bool("shuffle") {
		model.IsShuffled = true
		model.Cards = encard.Shuffle(model.Cards)
	}

	program := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		return err
	}
	return nil
}
