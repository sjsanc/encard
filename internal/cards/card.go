package cards

type Card interface {
	Render() string
	Deck() string
	Flipped() bool
	Flip()
}
