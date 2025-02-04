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
		case " ":
			if !m.IsFlipped {
				m.IsFlipped = true
			} else {
				m.NextCard()
			}
		}
	}

	return m, nil
}
