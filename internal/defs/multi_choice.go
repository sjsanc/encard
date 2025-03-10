package defs

import (
	"strings"

	"github.com/sjsanc/encard/internal/styles"
)

type Choice struct {
	Text    string
	Correct bool
}

type MultiChoice struct {
	Base
	Choices []Choice
	Current int
}

func NewMultiChoice(deck string, front string, choiceDefs map[string]bool) *MultiChoice {
	choices := make([]Choice, 0, len(choiceDefs))
	for choice, correct := range choiceDefs {
		choices = append(choices, Choice{Text: choice, Correct: correct})
	}

	return &MultiChoice{
		Base: Base{
			variant: "multichoice",
			deck:    deck,
			front:   front,
		},
		Choices: choices,
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
			if c.Current == i && choice.Correct {
				sb.WriteString(styles.Correct.Faint(faint).Render("* "+choice.Text+" (correct!)") + "\n")
			} else if c.Current == i && !choice.Correct {
				sb.WriteString(styles.Incorrect.Faint(faint).Render("* "+choice.Text+" (incorrect!)") + "\n")
			} else if c.Current != i && choice.Correct {
				sb.WriteString(styles.IncorrectUnselected.Faint(faint).Render("- "+choice.Text+" (answer)") + "\n")
			} else {
				sb.WriteString("- " + choice.Text + "\n")
			}
		} else {
			if c.Current == i {
				sb.WriteString(styles.Selected.Faint(faint).Render("* "+choice.Text) + "\n")
			} else {
				sb.WriteString("- " + choice.Text + "\n")
			}
		}
	}

	return sb.String()
}
