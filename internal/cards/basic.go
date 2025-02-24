package cards

type Basic struct {
	Deck  string
	Front string
	Back  string
}

func NewBasic(deck string, front string, back string) *Basic {
	return &Basic{
		Deck:  deck,
		Front: front,
		Back:  back,
	}
}

func (c *Basic) Update(input string) bool {
	return true // flip the card
}
