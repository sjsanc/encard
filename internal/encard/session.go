package encard

import (
	"math/rand/v2"

	"github.com/sjsanc/encard/internal/defs"
)

// A Session contains the loaded cards and the current card being displayed
type Session struct {
	cards    []defs.Card
	current  int
	finished bool // Whether the session is finished
	opts     *Opts
}

type Opts struct {
	shuffled bool
	verbose  bool
}

var logger *Logger

func NewSession(cards []defs.Card, opts *Opts) *Session {
	if opts.verbose {
		logger = NewLogger(true)
		logger.Printf("verbose logging enabled")
	} else {
		logger = NewLogger(false)
	}

	logger.Printf("%d cards loaded", len(cards))

	if opts.shuffled {
		logger.Println("shuffling cards")
		shuffle(cards)
	}

	return &Session{
		cards: cards,
		opts:  opts,
	}
}

func (s *Session) Update(key string) {
	card := s.cards[s.current]

	if !card.Flipped() {
		card.Update(key)
	} else {
		// Have to "enter" to advance
		if key == "enter" {
			s.NextCard()
		}
	}

	if s.finished {
		return
	}
}

func (s *Session) CurrentCard() defs.Card {
	return s.cards[s.current]
}

func (s *Session) NextCard() {
	if s.current >= len(s.cards)-1 {
		s.finished = true
		return
	}
	s.current++
}

func shuffle(cards []defs.Card) {
	perm := rand.Perm(len(cards))
	for i, v := range perm {
		cards[i] = cards[v]
	}
}
