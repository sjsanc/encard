package encard

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Width        int
	Height       int
	Cards        []*Card
	CurrentIndex int
	IsFlipped    bool
	IsCompleted  bool
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) GetCurrentCard() *Card {
	return m.Cards[m.CurrentIndex]
}

func (m *Model) NextCard() {
	m.IsFlipped = false
	m.CurrentIndex++
	if m.CurrentIndex >= len(m.Cards) {
		m.IsCompleted = true
	}
}
