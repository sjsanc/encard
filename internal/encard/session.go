package encard

import (
	"math/rand/v2"

	"github.com/sjsanc/encard/internal/cards"
)

// A Session contains the loaded cards and the current card being displayed
type Session struct {
	cards    []cards.Card
	current  int
	finished bool
	shuffled bool
}

func NewSession(cards []cards.Card, shuffled bool) *Session {
	if shuffled {
		shuffle(cards)
	}

	return &Session{
		cards:    cards,
		shuffled: shuffled,
	}
}

func (s *Session) Update(key string) {
	card := s.cards[s.current]
	flipped := card.Update(key)
	if flipped {
		s.NextCard()
	}
}

func (s *Session) NextCard() {
	s.current++
	if s.current >= len(s.cards) {
		s.finished = true
	}
}

func shuffle(cards []cards.Card) {
	perm := rand.Perm(len(cards))
	for i, v := range perm {
		cards[i] = cards[v]
	}
}
