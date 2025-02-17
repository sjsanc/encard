package encard

import tea "github.com/charmbracelet/bubbletea"

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.Advance()
		case " ":
			if ma, ok := m.CurrentCard().(*MultipleAnswerCard); ok {
				ma.ToggleChoice()
			}
		case "down":
			if mc, ok := m.CurrentCard().(*MultipleChoiceCard); ok {
				mc.NextChoice()
			}
			if mc, ok := m.CurrentCard().(*MultipleAnswerCard); ok {
				mc.NextChoice()
			}
		case "up":
			if mc, ok := m.CurrentCard().(*MultipleChoiceCard); ok {
				mc.PrevChoice()
			}
			if mc, ok := m.CurrentCard().(*MultipleAnswerCard); ok {
				mc.PrevChoice()
			}
		}
	}

	return m, nil
}
