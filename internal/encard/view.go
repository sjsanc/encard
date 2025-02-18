package encard

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	lg "github.com/charmbracelet/lipgloss"
)

func (m *Model) View() string {
	sb := strings.Builder{}

	if m.IsCompleted {
		sb.WriteString("You've completed the deck!\n")
		return sb.String()
	}

	card := m.CurrentCard()

	quarter := m.Width / 4
	half := m.Width - quarter - quarter

	sb.WriteString(card.Render())

	base := lg.NewStyle().
		Height(m.Height)

	left := base.
		Width(quarter).
		Align(lg.Right, lg.Top).
		Render(card.Deck())

	mid := base.
		Width(half).
		Align(lg.Left, lg.Top).
		Padding(0, 2).
		Render(sb.String())

	right := base.
		Width(quarter).
		Render("")

	block := lg.JoinHorizontal(
		lipgloss.Top,
		left,
		mid,
		right,
	)

	return block
}

// func (m *Model) buildLeftCol(w int, h int) string {
// 	block := lipgloss.NewStyle().Width(w).Height(h).Align(lipgloss.Right, lipgloss.Top)

// 	card := m.GetCurrentCard()

// 	body := strings.Builder{}

// 	body.WriteString(card.Deck + "\n")

// 	if m.IsShuffled {
// 		s := lipgloss.NewStyle().
// 			Foreground(lipgloss.Color("#39BAE6")).
// 			Bold(true).
// 			Render("shuffled")
// 		body.WriteString(s + "\n")
// 	}

// 	count := lipgloss.NewStyle().
// 		Foreground(lipgloss.Color("#FF8F40")).
// 		Render(fmt.Sprintf("%d of %d", m.CurrentIndex+1, len(m.Cards)))

// 	body.WriteString(count + "\n")

// 	cmd1 := lipgloss.NewStyle().
// 		Foreground(lipgloss.Color("#E6B673")).
// 		Render("press 'space' to flip")

// 	body.WriteString(cmd1 + "\n")

// 	cmd2 := lipgloss.NewStyle().
// 		Foreground(lipgloss.Color("#E6B673")).
// 		Render("press 'q' to quit")

// 	body.WriteString(cmd2 + "\n")

// 	return block.Render(body.String())
// }

// func renderCard(body *strings.Builder, card *Card, isFlipped bool, darken bool, width int) {
// 	color := "FF8F40"
// 	if darken {
// 		color = darkenHex(color, 0.5)
// 	}

// 	text := wordwrap.WrapString(card.Front, uint(width-2))

// 	front := lipgloss.NewStyle().
// 		Foreground(lipgloss.Color("#" + color)).
// 		Bold(true).
// 		Width(width).
// 		Render(text)

// 	body.WriteString(front + "\n")

// 	if isFlipped {
// 		color = "FFFFFF"
// 		if darken {
// 			color = darkenHex(color, 0.5)
// 		}

// 		text := wordwrap.WrapString(card.Back, uint(width-2))

// 		back := lipgloss.NewStyle().
// 			Foreground(lipgloss.Color("#" + color)).
// 			Width(width).
// 			Render(text)

// 		body.WriteString(back + "\n")
// 	}

// 	body.WriteString("\n")
// }

// func (m *Model) buildCenterCol(w int, h int) string {
// 	block := lipgloss.NewStyle().
// 		Width(w).
// 		Height(h).
// 		Padding(0, 1).
// 		Align(lipgloss.Left, lipgloss.Top)

// 	card := m.GetCurrentCard()

// 	body := &strings.Builder{}

// 	if m.IsCompleted {
// 		body.WriteString("You've completed the deck!\n")
// 		body.WriteString("\n")
// 	}

// 	renderCard(body, card, m.IsFlipped, false, w)

// 	if m.CurrentIndex > 0 {
// 		for i := m.CurrentIndex - 1; i >= 0; i-- {
// 			renderCard(body, m.Cards[i], true, true, w)
// 		}
// 	}

// 	lines := strings.Split(body.String(), "\n")
// 	if len(lines) > h {
// 		lines = lines[:h]
// 	}

// 	return block.Render(strings.Join(lines, "\n"))
// }

// func (m *Model) View() string {
// 	currentCard := m.GetCurrentCard()
// 	if currentCard == nil {
// 		return ""
// 	}

// 	sideWidth := 30
// 	centerWidth := m.Width - sideWidth

// 	left := m.buildLeftCol(sideWidth, m.Height)
// 	center := m.buildCenterCol(centerWidth, m.Height)

// 	block := lipgloss.JoinHorizontal(
// 		lipgloss.Bottom,
// 		left,
// 		center,
// 	)

// 	return block
// }
