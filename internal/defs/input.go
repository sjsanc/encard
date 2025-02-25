package defs

import (
	"math"

	"github.com/agnivade/levenshtein"
)

type Input struct {
	Base
	Deck    string
	Front   string
	Answer  string
	Input   string
	Matched bool
}

func NewInput(deck string, front string, answer string) *Input {
	return &Input{
		Deck:   deck,
		Front:  front,
		Answer: answer,
	}
}

func (c *Input) Update(key string) {
	switch key {
	case "backspace":
		if len(c.Input) > 0 {
			c.Input = c.Input[:len(c.Input)-1]
		}
	case "enter":
		c.Flip()
	default:
		c.Input += key
	}

	tolerance := int(math.Max(5.0, float64(len(c.Answer)/10))) // 10% of length, but at least 5
	c.Matched = levenshtein.ComputeDistance(c.Answer, c.Input) <= tolerance
}
