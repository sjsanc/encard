package defs

import (
	"strings"

	"github.com/sjsanc/encard/internal/styles"
)

type MultiChoice struct {
	Base
	Choices []string
	Answer  int
	Current int
}

func NewMultiChoice(deck string, front string, choices []string, answer int) *MultiChoice {
	return &MultiChoice{
		Base: Base{
			deck:  deck,
			front: front,
		},
		Choices: choices,
		Answer:  answer,
	}
}

func (c *MultiChoice) Update(key string) {
	switch key {
	case "up":
		c.Current = (c.Current - 1 + len(c.Choices)) % len(c.Choices)
	case "down":
		c.Current = (c.Current + 1) % len(c.Choices)
	case "enter":
		c.Flip()
	}
}

func (c *MultiChoice) Render(faint bool) string {
	sb := strings.Builder{}
	sb.WriteString(styles.Question.Faint(faint).Render(c.front) + "\n")

	for i, choice := range c.Choices {
		if c.flipped {
			if c.Current == i && c.Answer == i {
				sb.WriteString(styles.Correct.Faint(faint).Render("* "+choice+" (correct!)") + "\n")
			}

			if c.Current == i && c.Answer != i {
				sb.WriteString(styles.Incorrect.Faint(faint).Render("* "+choice+" (incorrect!)") + "\n")
			}

			if c.Current != i && c.Answer == i {
				sb.WriteString(styles.IncorrectUnselected.Faint(faint).Render("- "+choice+" (answer)") + "\n")
			}

			if c.Current != i && c.Answer != i {
				sb.WriteString("- " + choice + "\n")
			}

		} else {

			if c.Current == i {
				sb.WriteString(styles.Selected.Faint(faint).Render("* "+choice) + "\n")
			} else {
				sb.WriteString("- " + choice + "\n")
			}
		}
	}

	return sb.String()
}
