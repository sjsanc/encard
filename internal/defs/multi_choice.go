package defs

type MultiChoice struct {
	Base
	Choices []string
	Answer  int
	Current int
}

func NewMultiChoice(deck string, front string, choices []string, answer int) *MultiChoice {
	return &MultiChoice{
		Base: Base{
			deck:  deck,
			front: front,
		},
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

func (c *MultiChoice) Render(faint bool) string {
	return ""
}
