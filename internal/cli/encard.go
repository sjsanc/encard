package cli

import (
	"context"
	"fmt"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sjsanc/encard/internal/encard"
	"github.com/urfave/cli/v3"
)

// Shuffles a slice of Cards
// func Shuffle(cards []cards.Card) []cards.Card {
// 	shuffled := make([]cards.Card, len(cards))
// 	perm := rand.Perm(len(cards))
// 	for i, v := range perm {
// 		shuffled[v] = cards[i]
// 	}
// 	fmt.Println("Shuffled")
// 	return shuffled
// }

func encardAction(ctx context.Context, cmd *cli.Command) error {
	args := cmd.Args().Slice()

	_, err := encard.ResolveRootPath()
	if err != nil {
		return fmt.Errorf("unable to resolve a default Card directory: %w", err)
	}

	var cards []encard.Card

	for _, arg := range args {
		if filepath.IsAbs(arg) {
			deck, err := encard.LoadDeckFromPath(arg)
			if err != nil {
				return fmt.Errorf("%w", err)
			}
			cards = append(cards, deck...)
		}
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
