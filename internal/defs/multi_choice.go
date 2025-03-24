package defs

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
