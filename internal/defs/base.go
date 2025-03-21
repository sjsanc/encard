package defs

type Card interface {
	Variant() string   // card type
	Deck() string      // deck name
	Front() string     // front of card (i.e question)
	Flipped() bool     // whether the card has been flipped
	Flip()             // flip the card
	Update(key string) // update the card based on key
}

type Base struct {
	variant string
	deck    string
	front   string
	flipped bool
}

func (b *Base) Variant() string {
	return b.variant
}

func (b *Base) Deck() string {
	return b.deck
}

func (b *Base) Front() string {
	return b.front
}

func (b *Base) Flipped() bool {
	return b.flipped
}

func (b *Base) Flip() {
	b.flipped = true // can't unflip a card
}
