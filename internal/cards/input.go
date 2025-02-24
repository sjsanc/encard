package cards

import (
	"math"

	"github.com/agnivade/levenshtein"
)

type Input struct {
	Deck         string
	Front        string
	Answer       string
	CurrentInput string
	IsMatched    bool
}

func NewInput(deck string, front string, answer string) *Input {
	return &Input{
		Deck:   deck,
		Front:  front,
		Answer: answer,
	}
}

func (c *Input) Update(key string) bool {
	switch key {
	case "backspace":
		if len(c.CurrentInput) > 0 {
			c.CurrentInput = c.CurrentInput[:len(c.CurrentInput)-1]
		}
	case "enter":
		return true // flip the card
	default:
		c.CurrentInput += key
	}

	tolerance := int(math.Max(5.0, float64(len(c.Answer)/10))) // 10% of length, but at least 5
	c.IsMatched = levenshtein.ComputeDistance(c.Answer, c.CurrentInput) <= tolerance

	return false
}
