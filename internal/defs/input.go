package defs

import (
	"math"

	"github.com/agnivade/levenshtein"
	s "github.com/sjsanc/encard/internal/styles"
)

type Input struct {
	Base
	answer  string
	input   string
	matched bool
}

func NewInput(deck string, front string, answer string) *Input {
	return &Input{
		Base: Base{
			deck:  deck,
			front: front,
		},
		answer: answer,
	}
}

func (c *Input) Update(key string) {
	switch key {
	case "backspace":
		if len(c.input) > 0 {
			c.input = c.input[:len(c.input)-1]
		}
	case "enter":
		c.Flip()
	default:
		c.input += key
	}

	tolerance := int(math.Max(5.0, float64(len(c.answer)/10))) // 10% of length, but at least 5
	c.matched = levenshtein.ComputeDistance(c.answer, c.input) <= tolerance
}

func (c *Input) Render(faint bool) string {
	sb := s.Question.Faint(faint).Render(c.front) + "\n"

	if c.flipped {
		if c.matched {
			sb += s.Correct.Faint(faint).Render(c.input) + "\n"
		} else {
			sb += s.Incorrect.Faint(faint).Render(c.input) + "\n"
		}
		sb += s.Selected.Faint(faint).Render(c.answer) + "\n"
	}

	return sb
}
