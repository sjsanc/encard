package defs

type MultiChoice struct {
	Base
	Deck    string
	Front   string
	Choices []string
	Answer  int
	Current int
}

func NewMultiChoice(deck string, front string, choices []string, answer int) *MultiChoice {
	return &MultiChoice{
		Deck:    deck,
		Front:   front,
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
