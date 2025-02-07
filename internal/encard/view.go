package encard

import (
	"fmt"
	"slices"
	"strconv"
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

func clamp(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

func darkenHex(hex string, factor float64) string {
	// Parse the hex string as individual R, G, B components
	r, _ := strconv.ParseInt(hex[0:2], 16, 0)
	g, _ := strconv.ParseInt(hex[2:4], 16, 0)
	b, _ := strconv.ParseInt(hex[4:6], 16, 0)

	// Apply the darkening factor
	newR := int(float64(r) * factor)
	newG := int(float64(g) * factor)
	newB := int(float64(b) * factor)

	// Ensure values stay within [0, 255]
	newR = clamp(newR, 0, 255)
	newG = clamp(newG, 0, 255)
	newB = clamp(newB, 0, 255)

	return fmt.Sprintf("%02X%02X%02X", newR, newG, newB)
}

func (m *Model) buildCenterCol(w int, h int) string {
	card := m.GetCurrentCard()
	body := strings.Builder{}
	block := lipgloss.NewStyle().Width(w).Height(h).Padding(0, 1).Align(lipgloss.Left, lipgloss.Bottom)

	// TODO: Add reverse history
	count := h / 3
	history := make([]*Card, 0, count)
	for i, card := range m.Cards {
		if i > m.CurrentIndex-count && i < m.CurrentIndex {
			history = append(history, card)
		}
	}

	baseColor := "7FD962"
	shades := []string{}
	for i := 0; i < count; i++ {
		factor := 1 - float64(i)/float64(count)
		shades = append(shades, darkenHex(baseColor, factor))
	}

	shades = shades[:len(history)]
	slices.Reverse(shades)

	for i, card := range history {
		color := shades[i]
		body.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#" + color)).Width(w).Bold(true).Render(card.Front))
		body.WriteString(card.Back)
		body.WriteString("\n\n")
	}

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
