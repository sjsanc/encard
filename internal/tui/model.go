package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sjsanc/encard/internal/encard"
)

type Model struct {
	width   int
	height  int
	session *encard.Session
}

func NewModel(session *encard.Session) *Model {
	return &Model{
		session: session,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}
