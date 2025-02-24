package cards

type MultiChoice struct {
	Deck          string
	Front         string
	Choices       []string
	Answer        int
	CurrentChoice int
}

func NewMultiChoice(deck string, front string, choices []string, answer int) *MultiChoice {
	return &MultiChoice{
		Deck:    deck,
		Front:   front,
		Choices: choices,
		Answer:  answer,
	}
}

func (c *MultiChoice) Update(key string) bool {
	switch key {
	case "up":
		c.CurrentChoice = (c.CurrentChoice - 1 + len(c.Choices)) % len(c.Choices)
	case "down":
		c.CurrentChoice = (c.CurrentChoice + 1) % len(c.Choices)
	case "space":
		c.Answer = c.CurrentChoice
		return true
	}
	return false
}
