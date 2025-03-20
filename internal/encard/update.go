package encard

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "q", "ctrl+c":
			return m, tea.Quit
		default:
			m.session.Update(key)

			// ClearScreen allows kitty image to be cleared
			// TODO: save the images in cache so that we can apply a "dark" filter to it with Image
			// Also means we can have relative paths
			// TODO: add jpeg support (requires more complex flags)
			return m, tea.ClearScreen
		}
	}

	return m, nil
}
