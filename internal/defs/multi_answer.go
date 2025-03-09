package defs

import (
	"slices"
	"strings"

	s "github.com/sjsanc/encard/internal/styles"
)

type MultiAnswer struct {
	Base
	Choices  []string
	Answers  []int
	Selected []int
	Current  int
}

func NewMultiAnswer(deck string, front string, choices []string, answers []int) *MultiAnswer {
	return &MultiAnswer{
		Base: Base{
			variant: "multianswer",
			deck:    deck,
			front:   front,
		},
		Choices:  choices,
		Answers:  answers,
		Selected: make([]int, 0),
	}
}

func (c *MultiAnswer) Update(key string) {
	switch key {
	case "up":
		c.Current = (c.Current - 1 + len(c.Choices)) % len(c.Choices)
	case "down":
		c.Current = (c.Current + 1) % len(c.Choices)
	case " ":
		// Toggle selection of the current choice
		index := -1
		for i, choice := range c.Selected {
			if choice == c.Current {
				index = i
				break
			}
		}
		if index == -1 {
			c.Selected = append(c.Selected, c.Current)
		} else {
			c.Selected = append(c.Selected[:index], c.Selected[index+1:]...)
		}
	case "enter":
		c.Flip()
	}
}

func (c *MultiAnswer) Render(faint bool) string {
	sb := strings.Builder{}
	sb.WriteString(s.Question.Faint(faint).Render(c.front) + "\n")

	for i, choice := range c.Choices {
		if c.flipped {
			// Selected + Correct
			if slices.Contains(c.Selected, i) && slices.Contains(c.Answers, i) {
				sb.WriteString(s.Correct.Faint(faint).Render("[x] "+choice+" (correct!)") + "\n")
			}

			// Selected + Incorrect
			if slices.Contains(c.Selected, i) && !slices.Contains(c.Answers, i) {
				sb.WriteString(s.Incorrect.Faint(faint).Render("[x] "+choice+" (incorrect!)") + "\n")
			}

			// Not Selected + Correct
			if !slices.Contains(c.Selected, i) && slices.Contains(c.Answers, i) {
				sb.WriteString(s.IncorrectUnselected.Faint(faint).Render("[ ] "+choice+" (answer)") + "\n")
			}

			// Not Selected + Incorrect
			if !slices.Contains(c.Selected, i) && !slices.Contains(c.Answers, i) {
				sb.WriteString(s.Base.Faint(faint).Render("[ ] "+choice) + "\n")
			}
		} else {
			if slices.Contains(c.Selected, i) && c.Current == i {
				sb.WriteString(s.Selected.Faint(faint).Render("[x] "+choice) + "\n")
			} else if slices.Contains(c.Selected, i) {
				sb.WriteString(s.Selected.Faint(faint).Render("[x] "+choice) + "\n")
			} else if c.Current == i {
				sb.WriteString(s.Selected.Faint(faint).Render("[ ] "+choice) + "\n")
			} else {
				sb.WriteString(s.Base.Faint(faint).Render("[ ] "+choice) + "\n")
			}
		}
	}

	return sb.String()
}
