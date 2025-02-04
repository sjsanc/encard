package encard

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var MIN_WIDTH = 60

func getWidths(max int) (int, int) {
	centerWidth := max / 4
	if centerWidth < MIN_WIDTH {
		centerWidth = MIN_WIDTH
	}
	remainingWidth := max - centerWidth
	sideWidth := remainingWidth / 2
	return sideWidth, centerWidth
}

func (m *Model) buildLeftCol(w int, h int) string {
	card := m.GetCurrentCard()
	body := strings.Builder{}
	block := lipgloss.NewStyle().Width(w).Height(h).Padding(0, 1).Align(lipgloss.Right, lipgloss.Bottom)

	body.WriteString(card.Deck + "\n")

	count := fmt.Sprintf("%d of %d", m.CurrentIndex+1, len(m.Cards))
	body.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#FF8F40")).Render(count + "\n"))

	body.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#30384A")).Render("press 'space' to flip"))

	return block.Render(body.String())
}

func (m *Model) buildCenterCol(w int, h int) string {
	card := m.GetCurrentCard()
	body := strings.Builder{}
	block := lipgloss.NewStyle().Width(w).Height(h).Padding(0, 1).Align(lipgloss.Left, lipgloss.Bottom)

	// TODO: Add reverse history

	body.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#7FD962")).Width(w).Bold(true).Render(card.Front))

	if m.IsFlipped {
		body.WriteString(card.Back)
	} else {
		body.WriteString("\n")
	}

	return block.Render(body.String())
}

func (m *Model) buildRightCol(w int, h int) string {
	block := lipgloss.NewStyle().Width(w).Height(h).Padding(0, 1).Align(lipgloss.Left, lipgloss.Bottom)
	return block.Render("Press 'q' to quit")
}

func (m *Model) View() string {
	currentCard := m.GetCurrentCard()
	if currentCard == nil {
		return ""
	}

	sideWidth, centerWidth := getWidths(m.Width)

	left := m.buildLeftCol(sideWidth, m.Height/2)
	center := m.buildCenterCol(centerWidth, m.Height/2)
	right := m.buildRightCol(sideWidth, m.Height/2)

	block := lipgloss.JoinHorizontal(
		lipgloss.Bottom,
		left,
		center,
		right,
	)

	return block
}
