package encard

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m *Model) View() string {
	out := strings.Builder{}

	if m.IsCompleted {
		block := lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, "You're done!")
		return block
	}

	currentCard := m.GetCurrentCard()
	if currentCard == nil {
		return ""
	}

	deckName := lipgloss.NewStyle().Width(m.Width / 4).Align(lipgloss.Left)
	out.WriteString(deckName.Render(currentCard.Deck))

	out.WriteString("\n")

	countStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF8F40")).Width(m.Width / 4).Align(lipgloss.Left)
	count := fmt.Sprintf("%d/%d", m.CurrentIndex+1, len(m.Cards))
	out.WriteString(countStyle.Render(count))

	out.WriteString("\n\n")

	frontStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#7FD962")).Width(m.Width / 4).Bold(true).Align(lipgloss.Left)
	out.WriteString(frontStyle.Render(currentCard.Front))

	out.WriteString("\n")

	back := lipgloss.NewStyle().Align(lipgloss.Left).Width(m.Width / 4)

	if m.IsFlipped {
		out.WriteString(back.Render(currentCard.Back))
	} else {
		out.WriteString(" ")
	}

	block := lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, out.String())

	return block
}
