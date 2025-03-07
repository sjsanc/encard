package encard

import (
	lg "github.com/charmbracelet/lipgloss"
)

var ns = lg.NewStyle()

func (m *Model) renderLeft(w int) string {
	line := ns.Width(w).Padding(0, 2).Align(lg.Right)

	decks := make([]string, 0, len(m.session.decks))
	for _, deck := range m.session.decks {
		current := false
		if deck == m.session.CurrentCard().Deck() {
			current = true
		}
		prefix := ""
		if current {
			prefix = "> "
		}
		decks = append(decks, line.Faint(!current).Render(prefix+deck))
	}

	return lg.JoinVertical(
		lg.Top,
		decks...,
	)
}

func (m *Model) renderMid(w int) string {
	s := ns.Width(w)

	card := m.session.CurrentCard()

	block := lg.JoinVertical(
		lg.Top,
		s.Render(card.Render(false))+"\n",
	)

	history := m.session.History()
	for _, h := range history {
		block = lg.JoinVertical(
			lg.Top,
			block,
			s.Render(h.Render(true))+"\n",
		)
	}

	return block
}

func (m *Model) View() string {

	leftW := m.width / 4
	midW := m.width - leftW
	if midW > 80 {
		midW = 80
	}

	return lg.JoinHorizontal(
		lg.Top,
		m.renderLeft(leftW),
		m.renderMid(midW),
	)
}
