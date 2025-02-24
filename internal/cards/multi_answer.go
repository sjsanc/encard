package cards

type MultiAnswer struct {
	Deck            string
	Front           string
	Choices         []string
	Answers         []int
	SelectedChoices []int
	CurrentChoice   int
}

func NewMultiAnswer(deck string, front string, choices []string, answers []int) *MultiAnswer {
	return &MultiAnswer{
		Deck:            deck,
		Front:           front,
		Choices:         choices,
		Answers:         answers,
		SelectedChoices: make([]int, 0),
	}
}

func (c *MultiAnswer) Update(key string) bool {
	switch key {
	case "up":
		c.CurrentChoice = (c.CurrentChoice - 1 + len(c.Choices)) % len(c.Choices)
	case "down":
		c.CurrentChoice = (c.CurrentChoice + 1) % len(c.Choices)
	case "space":
		// Toggle selection of the current choice
		index := -1
		for i, choice := range c.SelectedChoices {
			if choice == c.CurrentChoice {
				index = i
				break
			}
		}
		if index == -1 {
			c.SelectedChoices = append(c.SelectedChoices, c.CurrentChoice)
		} else {
			c.SelectedChoices = append(c.SelectedChoices[:index], c.SelectedChoices[index+1:]...)
		}
	case "enter":
		return true
	}
	return false
}
