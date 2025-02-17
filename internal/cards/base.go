package cards

type Card interface {
	Render() string
	Deck() string
	Flipped() bool
	Flip()
}

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
