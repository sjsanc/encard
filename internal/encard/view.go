package encard

import (
	"fmt"

	lg "github.com/charmbracelet/lipgloss"
	"github.com/sjsanc/encard/internal/styles"
)

var ns = lg.NewStyle()

func (m *Model) renderLeft(w int) string {
	base := ns.Width(w).Padding(0, 2).Align(lg.Right)

	lines := make([]string, 0, len(m.session.decks))

	if m.session.opts.shuffled {
		lines = append(lines, base.Inherit(styles.Selected).Bold(true).Render("shuffled"))
	}

	for _, deck := range m.session.decks {
		current := false
		if deck == m.session.CurrentCard().Deck() {
			current = true
		}
		prefix := ""
		if current {
			prefix = "> "
		}
		lines = append(lines, base.Faint(!current).Render(prefix+deck))
	}

	return lg.JoinVertical(
		lg.Top,
		lines...,
	)
}

func (m *Model) renderMid(w int) string {
	base := ns.Width(w)

	card := m.session.CurrentCard()

	block := lg.JoinVertical(
		lg.Top,
		base.Render(card.Render(false))+"\n",
	)

	history := m.session.History()
	for _, h := range history {
		block = lg.JoinVertical(
			lg.Top,
			block,
			base.Render(h.Render(true))+"\n",
		)
	}

	if m.session.Finished() {
		block = lg.JoinVertical(
			lg.Top,
			base.Render("Session finished! Press 'q' to exit."+"\n"),
			block,
		)
	}

	return block
}

func (m *Model) renderRight(w int) string {
	base := ns.Width(w).Padding(0, 2).Align(lg.Left)

	card := m.session.CurrentCard()

	block := lg.JoinVertical(
		lg.Top,
		base.Faint(true).Render(card.Variant()+" card"),
	)

	return block
}

func (m *Model) View() string {

	fmt.Printf("\033_Ga=d\033\\")

	leftW := m.width / 4
	midW := m.width - leftW - leftW
	if midW > 80 {
		midW = 80
	}

	return lg.JoinHorizontal(
		lg.Top,
		m.renderLeft(leftW),
		m.renderMid(midW),
		m.renderRight(leftW),
	)
}
