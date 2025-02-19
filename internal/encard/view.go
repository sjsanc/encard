package encard

func (m *Model) View() string {
	flex := Flex{
		Children: []Element{
			NewElement("ENCARD", "#EC4899", "", false),
			NewElement("", "#202226", "", true),
			NewElement("Hello", "#A855F7", "", false),
		},
	}

	return flex.Render(m.Width)
}
