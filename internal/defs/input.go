package defs

import (
	"math"
	"strings"

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
			variant: "input",
			deck:    deck,
			front:   front,
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
	sb := strings.Builder{}
	sb.WriteString(s.Question.Faint(faint).Render(c.front) + "\n")

	if c.flipped {
		sb.WriteString(s.Base.Render(c.input) + "\n")
		if c.matched {
			sb.WriteString(s.Correct.Render(c.answer) + "\n")
		} else {
			sb.WriteString(s.Incorrect.Render(c.answer) + "\n")
		}
	} else {
		sb.WriteString(s.Base.Render(c.input) + "\n")
	}

	return sb.String()
}
