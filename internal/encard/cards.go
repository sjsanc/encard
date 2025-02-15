package encard

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/sjsanc/encard/internal/styles"
)

type Flippable struct {
	flipped bool
}

func (f *Flippable) Flipped() bool {
	return f.flipped
}
func (f *Flippable) Flip() {
	f.flipped = !f.flipped
}

type Card interface {
	Render() string
	Deck() string
	Flipped() bool
	Flip()
}

// ================================================================================
// ### BASIC CARD
// ================================================================================

type BasicCard struct {
	Flippable
	deck  string
	front string
	back  string
}

func (c *BasicCard) Render() string {
	sb := strings.Builder{}
	sb.WriteString(styles.Question.Render(c.front) + "\n")
	if c.flipped {
		sb.WriteString(c.back)
	}
	return sb.String()
}

func (c *BasicCard) Deck() string {
	return c.deck
}

// ================================================================================
// ### MULTIPLE CHOICE CARD
// ================================================================================

type MultipleChoiceCard struct {
	Flippable
	deck          string
	Front         string
	Choices       []string
	Answer        int
	CurrentChoice int
}

func (c *MultipleChoiceCard) Render() string {
	sb := strings.Builder{}
	sb.WriteString(styles.Question.Render(c.Front) + "\n")

	for i, choice := range c.Choices {
		prefix := "- "
		if i == c.CurrentChoice {
			prefix = "* "
		}

		if c.Flipped() {
			switch {
			case i == c.Answer:
				sb.WriteString(styles.CorrectChoice.Render(prefix+choice+" (correct!)") + "\n")

			case i == c.CurrentChoice && i != c.Answer:
				sb.WriteString(styles.IncorrectChoice.Render(prefix+choice+" (incorrect)") + "\n")

			default:
				sb.WriteString(prefix + choice + "\n")
			}
		} else {
			if i == c.CurrentChoice {
				sb.WriteString(styles.CurrentChoice.Render(prefix+choice) + "\n")
			} else {
				sb.WriteString(prefix + choice + "\n")
			}
		}
	}

	return sb.String()
}

func (c *MultipleChoiceCard) Deck() string {
	return c.deck
}

func (c *MultipleChoiceCard) NextChoice() {
	if c.Flipped() {
		return
	}
	c.CurrentChoice = (c.CurrentChoice + 1) % len(c.Choices)
}

func (c *MultipleChoiceCard) PrevChoice() {
	if c.Flipped() {
		return
	}
	c.CurrentChoice = (c.CurrentChoice - 1 + len(c.Choices)) % len(c.Choices)
}

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
