package defs

type Answer struct {
	Text     string
	Correct  bool
	Selected bool
}

type MultiAnswer struct {
	Base
	Answers []Answer
	Current int
}

func NewMultiAnswer(deck string, front string, answerDefs map[string]bool) *MultiAnswer {
	answers := make([]Answer, 0, len(answerDefs))
	for answer, correct := range answerDefs {
		answers = append(answers, Answer{Text: answer, Correct: correct})
	}

	return &MultiAnswer{
		Base: Base{
			variant: "multianswer",
			deck:    deck,
			front:   front,
		},
		Answers: answers,
	}
}

func (c *MultiAnswer) Update(key string) {
	switch key {
	case "up":
		c.Current = (c.Current - 1 + len(c.Answers)) % len(c.Answers)
	case "down":
		c.Current = (c.Current + 1) % len(c.Answers)
	case " ":
		choice := c.Answers[c.Current]
		if choice.Selected {
			c.Answers[c.Current].Selected = false
		} else {
			c.Answers[c.Current].Selected = true
		}
	case "enter":
		c.Flip()
	}
}
