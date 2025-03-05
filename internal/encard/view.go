package encard

import "fmt"

func (m *Model) View() string {
	count := len(m.session.cards)

	return fmt.Sprintf("You have %d cards in your collection", count)
}
