package defs

import (
	"strings"

	s "github.com/sjsanc/encard/internal/styles"
)

type Basic struct {
	Base
	back string
}

func NewBasic(deck string, front string, back string) *Basic {
	return &Basic{
		Base: Base{
			deck:  deck,
			front: front,
		},
		back: back,
	}
}

func (c *Basic) Update(key string) {
	switch key {
	case "enter":
		c.Flip()
	}
}

func (c *Basic) Render(faint bool) string {
	sb := strings.Builder{}
	sb.WriteString(s.Question.Faint(faint).Render(c.front) + "\n")

	if c.flipped {
		sb.WriteString(s.Base.Faint(faint).Render(c.back))
	}

	return sb.String()
}
