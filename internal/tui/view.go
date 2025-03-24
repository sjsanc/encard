package tui

import (
	lg "github.com/charmbracelet/lipgloss"
	"github.com/sjsanc/encard/internal/styles"
)

var ns = lg.NewStyle()

func (m *Model) renderLeft(w int) string {
	base := ns.Width(w).Padding(0, 2).Align(lg.Right)

	lines := make([]string, 0, len(m.session.DeckNames))

	if m.session.Opts.Shuffled {
		lines = append(lines, base.Inherit(styles.Selected).Bold(true).Render("shuffled"))
	}

	for _, deck := range m.session.DeckNames {
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

// To render the middle column of Cards, we need to be aware of the terminal height.
// Each rendered card consists of text lines and images.
// The total height of a rendered card is the sum of newlines and the image height, rounded to the pixel line height.
// So we must render each card, determine its height, and subtract that from the terminal height.
// When the remaining height would be less than the height of the next card, we must:
// - truncate the card to the nearest newline
// - and, if the card has an image, re-render the truncated form of the card.
// We should store a list of these height-aware renderables on the Bubbletea model as this isn't session data.
// This renderable will wrap the Card interface.

func (m *Model) renderMid(w int) string {
	base := ns.Width(w)

	card := m.session.CurrentCard()

	block := lg.JoinVertical(
		lg.Top,
		base.Render(displayCard(card, false))+"\n",
	)

	history := m.session.History()

	for _, h := range history {
		block = lg.JoinVertical(
			lg.Top,
			block,
			base.Render(displayCard(h, true))+"\n",
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
