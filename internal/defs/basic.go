package defs

import (
	"strings"

	"github.com/sjsanc/encard/internal/image"
	s "github.com/sjsanc/encard/internal/styles"
)

type Basic struct {
	Base
	back []string
}

func NewBasic(deck string, front string, back []string) *Basic {
	return &Basic{
		Base: Base{
			variant: "basic",
			deck:    deck,
			front:   front,
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
		for _, line := range c.back {
			if strings.HasPrefix(line, "[](") {
				filepath := strings.TrimSuffix(strings.TrimPrefix(line, "[]("), ")")
				img := image.NewImage(filepath)
				sb.WriteString(img.Print() + "\n")
			} else {
				sb.WriteString(s.Base.Faint(faint).Render(line) + "\n")
			}
		}
	}

	return sb.String()
}
