package encard

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sjsanc/encard/internal/cards"
)

type DeckMap map[string][]cards.Card

type Model struct {
	width   int
	height  int
	session *Session
}

func NewModel(session *Session) *Model {
	return &Model{
		session: session,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}
