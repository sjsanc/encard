package encard

// Shuffles a slice of Cards
// func Shuffle(cards []Card) []Card {
// 	shuffled := make([]Card, len(cards))
// 	perm := rand.Perm(len(cards))
// 	for i, v := range perm {
// 		shuffled[v] = cards[i]
// 	}
// 	fmt.Println("Shuffled")
// 	return shuffled
// }

type Card interface {
	Render() string
	Deck() string
	Flipped() bool
	Flip()
}
