package cards

import (
	"strings"

	"github.com/agnivade/levenshtein"
	"github.com/sjsanc/encard/internal/styles"
)

type Input struct {
	CardBase
	answer string
	input  string
}

func NewInput(front string, answer string) *Input {
	return &Input{
		CardBase: CardBase{
			front: front,
		},
		answer: answer,
	}
}

func (c *Input) Render() string {
	sb := strings.Builder{}

	sb.WriteString(styles.Question.Render(c.front) + "\n")

	distance := levenshtein.ComputeDistance(c.input, c.answer)
	isCorrect := threshold(len(c.answer), distance)

	if c.flipped {
		if isCorrect {
			sb.WriteString(styles.Correct.Render(c.input) + "\n")
		} else {
			sb.WriteString(styles.Incorrect.Render(c.input) + "\n")
		}
		sb.WriteString(styles.Selected.Render(c.answer) + "\n")
	}

	return sb.String()
}

func threshold(length, distance int) bool {
	tolerance := max(5, length/10) // 10% of length, but at least 5
	return distance <= tolerance
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
