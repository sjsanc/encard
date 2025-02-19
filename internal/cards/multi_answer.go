package cards

import (
	"slices"
	"strings"

	"github.com/sjsanc/encard/internal/styles"
)

type MultiAnswer struct {
	CardBase
	choices  []string
	answers  []int
	selected []int
	current  int
}

func NewMultiAnswer(front string, choices []string, answers []int) *MultiAnswer {
	return &MultiAnswer{
		CardBase: CardBase{
			front: front,
		},
		choices:  choices,
		answers:  answers,
		selected: make([]int, 0),
	}
}

func (c *MultiAnswer) Render() string {
	sb := strings.Builder{}
	sb.WriteString(styles.Question.Render(c.front) + "\n")

	for i, choice := range c.choices {
		if c.flipped {
			// Selected + Correct
			if slices.Contains(c.selected, i) && slices.Contains(c.answers, i) {
				sb.WriteString(styles.Correct.Render("[x] "+choice+" (correct!)") + "\n")
			}

			// Selected + Incorrect
			if slices.Contains(c.selected, i) && !slices.Contains(c.answers, i) {
				sb.WriteString(styles.Incorrect.Render("[x] "+choice+" (incorrect!)") + "\n")
			}

			// Not Selected + Correct
			if !slices.Contains(c.selected, i) && slices.Contains(c.answers, i) {
				sb.WriteString(styles.IncorrectUnselected.Render("[ ] "+choice+" (answer)") + "\n")
			}

			// Not Selected + Incorrect
			if !slices.Contains(c.selected, i) && !slices.Contains(c.answers, i) {
				sb.WriteString("[ ] " + choice + "\n")
			}
		} else {
			if slices.Contains(c.selected, i) && c.current == i {
				sb.WriteString(styles.Selected.Render("[x] "+choice) + "\n")
			} else if slices.Contains(c.selected, i) {
				sb.WriteString("[x] " + choice + "\n")
			} else if c.current == i {
				sb.WriteString(styles.Selected.Render("[ ] "+choice) + "\n")
			} else {
				sb.WriteString("[ ] " + choice + "\n")
			}
		}
	}

	return sb.String()
}

func (c *MultiAnswer) NextChoice() {
	if c.flipped {
		return
	}
	c.current = (c.current + 1) % len(c.choices)
}

func (c *MultiAnswer) PrevChoice() {
	if c.flipped {
		return
	}
	c.current = (c.current - 1 + len(c.choices)) % len(c.choices)
}

func (c *MultiAnswer) ToggleChoice() {
	if slices.Contains(c.selected, c.current) {
		i := slices.Index(c.selected, c.current)
		c.selected = append(c.selected[:i], c.selected[i+1:]...)
	} else {
		c.selected = append(c.selected, c.current)
	}
}
