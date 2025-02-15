package encard

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Width            int
	Height           int
	Cards            []Card
	CurrentCardIndex int
	IsCompleted      bool
	IsShuffled       bool
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) CurrentCard() Card {
	return m.Cards[m.CurrentCardIndex]
}

func (m *Model) Advance() {
	current := m.CurrentCard()

	if !current.Flipped() {
		current.Flip()
		return
	}

	if m.CurrentCardIndex >= len(m.Cards)-1 {
		m.IsCompleted = true
		return
	}

	m.CurrentCardIndex++
}
