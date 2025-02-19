package encard

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sjsanc/encard/internal/cards"
)

type DeckMap map[string][]cards.Card

type Model struct {
	Width            int
	Height           int
	DeckMap          DeckMap
	CurrentDeckKey   string
	CurrentCardIndex int
	IsCompleted      bool
	IsShuffled       bool
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) CurrentDeck() []cards.Card {
	return m.DeckMap[m.CurrentDeckKey]
}

func (m *Model) CurrentCard() cards.Card {
	deck := m.CurrentDeck()
	return deck[m.CurrentCardIndex]
}

func (m *Model) Advance() {
	cards := m.CurrentDeck()
	current := m.CurrentCard()

	if !current.Flipped() {
		current.Flip()
		return
	}

	if m.CurrentCardIndex >= len(cards)-1 {
		m.IsCompleted = true
		return
	}

	m.CurrentCardIndex++
}
