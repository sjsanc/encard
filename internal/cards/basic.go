package cards

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/sjsanc/encard/internal/styles"
)

type Basic struct {
	CardBase
	back string
}

func NewBasic(front string, back string) *Basic {
	return &Basic{
		CardBase: CardBase{
			front: front,
		},
		back: back,
	}
}

func (c *Basic) Render() string {
	sb := strings.Builder{}

	sb.WriteString(styles.Question.Render(c.front) + "\n")

	if c.flipped {
		sb.WriteString(c.back)
	}

	result := lipgloss.NewStyle().
		Padding(0, 1).
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.Color("#542b7c")).
		Render(sb.String())

	return result
}
