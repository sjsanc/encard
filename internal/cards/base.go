package cards

type Card interface {
	Render() string
	Flipped() bool
	Flip()
}

type CardBase struct {
	front   string
	flipped bool
}

func (c *CardBase) Flipped() bool {
	return c.flipped
}

func (c *CardBase) Flip() {
	c.flipped = !c.flipped
}
