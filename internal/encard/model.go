package encard

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sjsanc/encard/internal/cards"
)

type Model struct {
	Width            int
	Height           int
	Cards            map[string][]cards.Card
	CurrentDeckKey   string
	CurrentCardIndex int
	IsCompleted      bool
	IsShuffled       bool
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) CurrentDeck() []cards.Card {
	return m.Cards[m.CurrentDeckKey]
}

func (m *Model) CurrentCard() cards.Card {
	deck := m.CurrentDeck()
	return deck[m.CurrentCardIndex]
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
