package encard

import (
	"fmt"
	"math/rand"
	"strings"
)

type Card interface {
	Render() string
}

type Flippable struct {
	IsFlipped bool
}

func (f *Flippable) Flip() {
	f.IsFlipped = !f.IsFlipped
}
func (f *Flippable) Unflip() {
	f.IsFlipped = false
}

type BasicCard struct {
	Flippable
	Deck      string
	Front     string
	Back      string
	IsFlipped bool
}

func (c *BasicCard) Render() string {
	sb := strings.Builder{}
	sb.WriteString(c.Front + "\n")
	if c.IsFlipped {
		sb.WriteString(c.Back)
	}
	return sb.String()
}

type MultipleChoiceCard struct {
	Flippable
	Deck      string
	Front     string
	Choices   []string
	Answer    int
	Selected  int
	IsFlipped bool
}

func (c *MultipleChoiceCard) Render() string {
	sb := strings.Builder{}
	sb.WriteString(c.Front + "\n")
	for i, choice := range c.Choices {
		// TODO: if flipped, just set colors
		if i == c.Selected {
			sb.WriteString(fmt.Sprintf("* %s\n", choice))
		} else {
			sb.WriteString(fmt.Sprintf("  %s\n", choice))
		}
	}
	return sb.String()
}

// Converts raw text into a slice of Cards
// func ParseCards(data, deckName string) []*Card {
// 	raw := strings.Split(data, "---")
// 	cards := make([]*Card, 0)

// 	for _, r := range raw {
// 		lines := strings.Split(strings.TrimSpace(r), "\n")
// 		if len(lines) < 2 {
// 			continue
// 		}
// 		front := strings.TrimPrefix(lines[0], "# ")
// 		back := lines[1]
// 		cards = append(cards, &Card{
// 			Deck:  deckName,
// 			Front: front,
// 			Back:  back,
// 		})
// 	}

// 	return cards
// }

// Parses all cards from a given path
// func ParseCardsFromPath(path string) ([]*Card, error) {
// 	var allCards []*Card
// 	err := filepath.Walk(path, func(entryPath string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}

// 		if info.IsDir() {
// 			return nil
// 		}

// 		if strings.HasSuffix(info.Name(), ".md") {
// 			deckName := strings.TrimSuffix(filepath.ToSlash(entryPath), ".md")
// 			data, err := os.ReadFile(entryPath)
// 			if err != nil {
// 				return err
// 			}

// 			allCards = append(allCards, ParseCards(string(data), deckName)...)
// 		}
// 		return nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}
// 	return allCards, nil
// }

// Shuffles a slice of Cards
func Shuffle(cards []Card) []Card {
	shuffled := make([]Card, len(cards))
	perm := rand.Perm(len(cards))
	for i, v := range perm {
		shuffled[v] = cards[i]
	}
	fmt.Println("Shuffled")
	return shuffled
}
