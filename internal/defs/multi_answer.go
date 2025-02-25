package defs

type MultiAnswer struct {
	Base
	Deck     string
	Front    string
	Choices  []string
	Answers  []int
	Selected []int
	Current  int
}

func NewMultiAnswer(deck string, front string, choices []string, answers []int) *MultiAnswer {
	return &MultiAnswer{
		Deck:     deck,
		Front:    front,
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
	case "space":
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
