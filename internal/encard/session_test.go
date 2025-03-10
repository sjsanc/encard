package encard

import (
	"path/filepath"
	"testing"

	"github.com/sjsanc/encard/internal/defs"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T, path string) *Session {
	testdataDir, err := filepath.Abs("testdata/session")
	if err != nil {
		t.Fatalf("failed to resolve testdata directory: %v", err)
	}
	cards, errors := LoadCards([]string{path}, testdataDir)
	assert.Empty(t, errors)
	assert.Equal(t, 2, len(cards))
	session := NewSession(cards, &Options{})
	assert.Equal(t, 0, session.current) // current card index is 0
	return session
}

func assertFinished(t *testing.T, session *Session) {
	assert.True(t, session.finished)
	assert.Equal(t, 1, session.current)
}

func TestSession_Basic(t *testing.T) {
	s := setup(t, "basic.md")

	s.Update("enter") // flip
	s.Update("enter") // next

	s.Update("enter") // flip
	s.Update("enter") // next

	assertFinished(t, s)
}

func TestSession_MultiAnswer(t *testing.T) {
	s := setup(t, "multianswer.md")

	s.Update(" ") // select
	card := s.CurrentCard().(*defs.MultiAnswer)
	assert.Equal(t, []int{0}, card.Selected)

	s.Update("down") // next
	card = s.CurrentCard().(*defs.MultiAnswer)
	assert.Equal(t, 1, card.Current)

	s.Update(" ") // select

	assert.Equal(t, []int{0, 1}, card.Selected)

	s.Update("down") // next
	card = s.CurrentCard().(*defs.MultiAnswer)
	assert.Equal(t, 2, card.Current)

	s.Update(" ") // select
	s.Update(" ") // de-select

	assert.Equal(t, []int{0, 1}, card.Selected)

	s.Update("enter") // flip
	s.Update("enter") // next

	s.Update("down")
	s.Update("down")
	s.Update("down")

	s.Update("enter") // flip
	s.Update("enter") // next

	assertFinished(t, s)

	card = s.CurrentCard().(*defs.MultiAnswer)
	assert.Empty(t, card.Selected)
}

func TestSession_MultiChoice(t *testing.T) {
	s := setup(t, "multichoice.md")

	assert.False(t, s.CurrentCard().Flipped())

	s.Update("enter") // flip
	card := s.CurrentCard().(*defs.MultiChoice)
	assert.Equal(t, 0, card.Current)

	s.Update("enter") // next
	s.Update("down")
	card = s.CurrentCard().(*defs.MultiChoice)
	assert.Equal(t, 1, card.Current)

	s.Update("enter") // flip
	assert.True(t, s.CurrentCard().Flipped())

	s.Update("enter") // next

	assertFinished(t, s)
}
