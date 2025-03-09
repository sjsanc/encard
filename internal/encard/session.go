package encard

import (
	"math/rand/v2"
	"slices"

	"github.com/sjsanc/encard/internal/defs"
)

// A Session contains the loaded cards and the current card being displayed
type Session struct {
	cards    []defs.Card
	decks    []string
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
		cards = shuffle(cards)
	}

	decks := make(map[string]bool)
	for _, card := range cards {
		decks[card.Deck()] = true
	}
	keys := make([]string, 0, len(decks))
	for k := range decks {
		keys = append(keys, k)
	}

	return &Session{
		cards: cards,
		opts:  opts,
		decks: keys,
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

func (s *Session) PrevCard() {
	if s.current == 0 {
		return
	}
	s.current--
}

func (s *Session) Finished() bool {
	if s.current == len(s.cards)-1 && s.cards[s.current].Flipped() {
		return true
	}
	return false
}

func (s *Session) History() []defs.Card {
	c := make([]defs.Card, s.current)
	copy(c, s.cards[:s.current])
	slices.Reverse(c)
	return c
}

func shuffle(cards []defs.Card) []defs.Card {
	shuffled := make([]defs.Card, len(cards))
	perm := rand.Perm(len(cards))
	for i, v := range perm {
		shuffled[v] = cards[i]
	}
	return shuffled
}
