package encard

import (
	"os"
	"path/filepath"
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
		front := strings.TrimPrefix(lines[0], "# ")
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
	err := parseCardsRecursive(path, path, &allCards)
	if err != nil {
		return nil, err
	}
	return allCards, nil
}

func parseCardsRecursive(rootPath, path string, cards *[]*Card) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		entryPath := path + "/" + entry.Name()

		if entry.IsDir() {
			if err := parseCardsRecursive(rootPath, entryPath, cards); err != nil {
				return err
			}
		} else if strings.HasSuffix(entry.Name(), ".md") {
			relativePath, err := filepath.Rel(rootPath, entryPath) // Trim the root dir path
			if err != nil {
				return err
			}

			deckName := strings.TrimSuffix(relativePath, ".md")
			data, err := os.ReadFile(entryPath)
			if err != nil {
				return err
			}

			*cards = append(*cards, ParseCards(string(data), deckName)...)
		}
	}

	return nil
}
