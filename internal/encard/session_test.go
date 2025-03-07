package encard

import (
	"testing"

	"github.com/sjsanc/encard/internal/defs"
	"github.com/stretchr/testify/assert"
)

func assertLoader(t *testing.T, c []defs.Card, e error) {
	assert.Nil(t, e)
	assert.Equal(t, 2, len(c))
}

func TestSession_Basic(t *testing.T) {
	cards, err := LoadCards("testdata/session/basic.md", nil, false)
	assertLoader(t, cards, err)

	session := NewSession(cards, &Opts{})

	session.Update("enter") // flip
	assert.True(t, session.CurrentCard().Flipped())

	session.Update("enter") // next
	assert.Equal(t, 1, session.current)
	assert.False(t, session.CurrentCard().Flipped())

	session.Update("enter") // flip
	assert.True(t, session.CurrentCard().Flipped())

	session.Update("enter") // next
	assert.True(t, session.finished)
	assert.Equal(t, 1, session.current)
}

func TestSession_MultiAnswer(t *testing.T) {
	cards, err := LoadCards("testdata/session/multianswer.md", nil, false)
	assertLoader(t, cards, err)

	session := NewSession(cards, &Opts{})

	session.Update("space") // select
	card := session.CurrentCard().(*defs.MultiAnswer)
	assert.Equal(t, []int{0}, card.Selected)

	session.Update("down") // next
	card = session.CurrentCard().(*defs.MultiAnswer)
	assert.Equal(t, 1, card.Current)

	session.Update("space") // select

	assert.Equal(t, []int{0, 1}, card.Selected)

	session.Update("down") // next

	assert.Equal(t, 2, card.Current)

	session.Update("space") // select
	session.Update("space") // de-select

	assert.Equal(t, []int{0, 1}, card.Selected)

	session.Update("enter") // flip

	assert.True(t, session.CurrentCard().Flipped())

	session.Update("enter") // next

	assert.False(t, session.CurrentCard().Flipped())

	session.Update("down")
	session.Update("down")
	session.Update("down")

	session.Update("enter") // flip

	assert.True(t, session.CurrentCard().Flipped())

	session.Update("enter") // next

	assert.True(t, session.finished)
	assert.Equal(t, 1, session.current)

	card = session.CurrentCard().(*defs.MultiAnswer)
	assert.Empty(t, card.Selected)
}

func TestSession_MultiChoice(t *testing.T) {
	cards, err := LoadCards("testdata/session/multichoice.md", nil, false)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(cards))

	session := NewSession(cards, &Opts{})

	assert.False(t, session.CurrentCard().Flipped())

	session.Update("enter") // flip
	card := session.CurrentCard().(*defs.MultiChoice)
	assert.Equal(t, 0, card.Current)

	session.Update("enter") // next
	session.Update("down")
	card = session.CurrentCard().(*defs.MultiChoice)
	assert.Equal(t, 1, card.Current)

	session.Update("enter") // flip
	assert.True(t, session.CurrentCard().Flipped())

	session.Update("enter") // next
	assert.True(t, session.finished)
}
