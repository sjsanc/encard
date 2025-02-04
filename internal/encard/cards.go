package encard

import (
	"os"
	"strings"
)

type Card struct {
	Deck  string
	Front string
	Back  string
}

func ParseCards(data, deckName string) []*Card {
	raw := strings.Split(data, "---")
	cards := make([]*Card, 0)

	for _, r := range raw {
		lines := strings.Split(strings.TrimSpace(r), "\n")
		if len(lines) < 2 {
			continue
		}
		front := lines[0]
		back := lines[1]
		cards = append(cards, &Card{
			Deck:  deckName,
			Front: front,
			Back:  back,
		})
	}

	return cards
}

func ParseCardsFromPath(path string) ([]*Card, error) {
	var allCards []*Card
	err := parseCardsRecursive(path, &allCards)
	if err != nil {
		return nil, err
	}
	return allCards, nil
}

func parseCardsRecursive(path string, cards *[]*Card) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		entryPath := path + "/" + entry.Name()

		if entry.IsDir() {
			if err := parseCardsRecursive(entryPath, cards); err != nil {
				return err
			}
		} else if strings.HasSuffix(entry.Name(), ".md") {
			data, err := os.ReadFile(entryPath)
			if err != nil {
				return err
			}
			*cards = append(*cards, ParseCards(string(data), entry.Name())...)
		}
	}

	return nil
}
