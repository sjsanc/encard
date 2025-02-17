package encard

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"

	"github.com/sjsanc/encard/internal/styles"
)

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
	flipped bool
	deck    string
	front   string
	back    string
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

func (c *BasicCard) Flipped() bool {
	return c.flipped
}

func (c *BasicCard) Flip() {
	c.flipped = !c.flipped
}

// ================================================================================
// ### MULTIPLE CHOICE CARD
// ================================================================================

type MultipleChoiceCard struct {
	flipped       bool
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

func (c *MultipleChoiceCard) Flipped() bool {
	return c.flipped
}

func (c *MultipleChoiceCard) Flip() {
	c.flipped = !c.flipped
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

// ================================================================================
// ### MULTIPLE ANSWER CARD
// ================================================================================

type MultipleAnswerCard struct {
	flipped       bool
	deck          string
	Front         string
	Choices       []string
	Answers       []int
	CurrentChoice int
	Selected      []int
}

func (c *MultipleAnswerCard) Render() string {
	sb := strings.Builder{}
	sb.WriteString(styles.Question.Render(c.Front) + "\n")

	for i, choice := range c.Choices {
		if c.Flipped() {
			// Selected + Correct
			if slices.Contains(c.Selected, i) && slices.Contains(c.Answers, i) {
				sb.WriteString(styles.CorrectChoice.Render("[x] "+choice+" (correct!)") + "\n")
			}

			// Selected + Incorrect
			if slices.Contains(c.Selected, i) && !slices.Contains(c.Answers, i) {
				sb.WriteString(styles.IncorrectChoice.Render("[x] "+choice+" (incorrect!)") + "\n")
			}

			// Not Selected + Correct
			if !slices.Contains(c.Selected, i) && slices.Contains(c.Answers, i) {
				sb.WriteString(styles.UnselectedChoice.Render("[ ] "+choice+" (answer)") + "\n")
			}

			// Not Selected + Incorrect
			if !slices.Contains(c.Selected, i) && !slices.Contains(c.Answers, i) {
				sb.WriteString("[ ] " + choice + "\n")
			}

		} else {
			prefix := "[ ] "
			if i == c.CurrentChoice {
				prefix = "[*] "
			}

			if slices.Contains(c.Selected, i) {
				prefix = "[x] "
			}

			sb.WriteString(prefix + choice + "\n")
		}
	}

	return sb.String()
}

func (c *MultipleAnswerCard) Deck() string {
	return c.deck
}

func (c *MultipleAnswerCard) Flipped() bool {
	return c.flipped
}

func (c *MultipleAnswerCard) Flip() {
	c.flipped = !c.flipped
}

func (c *MultipleAnswerCard) ToggleChoice() {
	if slices.Contains(c.Selected, c.CurrentChoice) {
		i := slices.Index(c.Selected, c.CurrentChoice)
		c.Selected = append(c.Selected[:i], c.Selected[i+1:]...)
	} else {
		c.Selected = append(c.Selected, c.CurrentChoice)
	}
}

func (c *MultipleAnswerCard) NextChoice() {
	if c.Flipped() {
		return
	}
	c.CurrentChoice = (c.CurrentChoice + 1) % len(c.Choices)
}

func (c *MultipleAnswerCard) PrevChoice() {
	if c.Flipped() {
		return
	}
	c.CurrentChoice = (c.CurrentChoice - 1 + len(c.Choices)) % len(c.Choices)
}
