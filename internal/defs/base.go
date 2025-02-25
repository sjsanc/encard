package defs

type Card interface {
	Update(key string)
	Flipped() bool
}

type Base struct {
	flipped bool
}

func (b *Base) Flipped() bool {
	return b.flipped
}

func (b *Base) Flip() {
	b.flipped = true // can't unflip a card
}
