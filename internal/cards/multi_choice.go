package cards

import (
	"strings"

	"github.com/sjsanc/encard/internal/styles"
)

type MultiChoice struct {
	CardBase
	choices []string
	answer  int
	current int
}

func NewMultiChoice(deck string, front string, choices []string, answer int) *MultiChoice {
	return &MultiChoice{
		CardBase: CardBase{
			deck:  deck,
			front: front,
		},
		choices: choices,
		answer:  answer,
	}
}

func (c *MultiChoice) Render() string {
	sb := strings.Builder{}
	sb.WriteString(styles.Question.Render(c.front) + "\n")

	for i, choice := range c.choices {
		if c.flipped {
			// Selected + Correct
			if c.current == i && c.answer == i {
				sb.WriteString(styles.Correct.Render("* "+choice+" (correct!)") + "\n")
			}

			// Selected + Incorrect
			if c.current == i && c.answer != i {
				sb.WriteString(styles.Incorrect.Render("* "+choice+" (incorrect!)") + "\n")
			}

			// Not Selected + Correct
			if c.current != i && c.answer == i {
				sb.WriteString(styles.IncorrectUnselected.Render("- "+choice+" (answer)") + "\n")
			}

			// Not Selected + Incorrect
			if c.current != i && c.answer != i {
				sb.WriteString("- " + choice + "\n")
			}
		} else {
			if c.current == i {
				sb.WriteString(styles.Selected.Render("* "+choice) + "\n")
			} else {
				sb.WriteString("- " + choice + "\n")
			}
		}
	}

	return sb.String()
}

func (c *MultiChoice) NextChoice() {
	if c.flipped {
		return
	}
	c.current = (c.current + 1) % len(c.choices)
}

func (c *MultiChoice) PrevChoice() {
	if c.flipped {
		return
	}
	c.current = (c.current - 1 + len(c.choices)) % len(c.choices)
}
