package defs

import (
	"strings"

	s "github.com/sjsanc/encard/internal/styles"
)

type Answer struct {
	Text     string
	Correct  bool
	Selected bool
}

type MultiAnswer struct {
	Base
	Answers []Answer
	Current int
}

func NewMultiAnswer(deck string, front string, answerDefs map[string]bool) *MultiAnswer {
	answers := make([]Answer, 0, len(answerDefs))
	for answer, correct := range answerDefs {
		answers = append(answers, Answer{Text: answer, Correct: correct})
	}

	return &MultiAnswer{
		Base: Base{
			variant: "multianswer",
			deck:    deck,
			front:   front,
		},
		Answers: answers,
	}
}

func (c *MultiAnswer) Update(key string) {
	switch key {
	case "up":
		c.Current = (c.Current - 1 + len(c.Answers)) % len(c.Answers)
	case "down":
		c.Current = (c.Current + 1) % len(c.Answers)
	case " ":
		choice := c.Answers[c.Current]
		if choice.Selected {
			c.Answers[c.Current].Selected = false
		} else {
			c.Answers[c.Current].Selected = true
		}
	case "enter":
		c.Flip()
	}
}

func (c *MultiAnswer) Render(faint bool) string {
	sb := strings.Builder{}
	sb.WriteString(s.Question.Faint(faint).Render(c.front) + "\n")

	for i, choice := range c.Answers {
		selected := choice.Selected
		correct := choice.Correct

		if c.flipped {
			if selected && correct {
				sb.WriteString(s.Correct.Faint(faint).Render("[x] "+choice.Text+" (correct!)") + "\n")
			} else if selected && !correct {
				sb.WriteString(s.Incorrect.Faint(faint).Render("[x] "+choice.Text+" (incorrect!)") + "\n")
			} else if !selected && correct {
				sb.WriteString(s.IncorrectUnselected.Faint(faint).Render("[ ] "+choice.Text+" (answer)") + "\n")
			} else {
				sb.WriteString(s.Base.Faint(faint).Render("[ ] "+choice.Text) + "\n")
			}
		} else {
			if c.Current == i {
				if selected {
					sb.WriteString(s.Selected.Render("[x] "+choice.Text) + "\n")
				} else {
					sb.WriteString(s.Selected.Render("[ ] "+choice.Text) + "\n")
				}
			} else {
				if selected {
					sb.WriteString(s.Base.Faint(faint).Render("[x] "+choice.Text) + "\n")
				} else {
					sb.WriteString(s.Base.Faint(faint).Render("[ ] "+choice.Text) + "\n")
				}
			}
		}
	}

	return sb.String()
}
