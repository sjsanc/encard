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
	IsShuffled   bool
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) GetCurrentCard() *Card {
	return m.Cards[m.CurrentIndex]
}

func (m *Model) NextCard() {
	if m.CurrentIndex >= len(m.Cards)-1 {
		m.IsCompleted = true
	} else {
		m.IsFlipped = false
		m.CurrentIndex++
	}
}
