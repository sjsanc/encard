package cards

type CardBase struct {
	front   string
	flipped bool
	deck    string
}

func (c *CardBase) Deck() string {
	return c.deck
}

func (c *CardBase) Flipped() bool {
	return c.flipped
}

func (c *CardBase) Flip() {
	c.flipped = !c.flipped
}
