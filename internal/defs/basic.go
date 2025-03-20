package defs

import (
	"fmt"
	"strings"

	s "github.com/sjsanc/encard/internal/styles"
)

type Basic struct {
	Base
	back string
	img  string
}

func NewBasic(deck string, front string, back string, img string) *Basic {
	return &Basic{
		Base: Base{
			variant: "basic",
			deck:    deck,
			front:   front,
		},
		back: back,
		img:  img,
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

	if c.img != "" {
		sb.WriteString(fmt.Sprintf("\033_Gi=1,t=f,q=1,f=100,a=T;%s\033\\", c.img))
		sb.WriteString("\n")
	}

	if c.flipped {
		sb.WriteString(s.Base.Faint(faint).Render(c.back))
	}

	return sb.String()
}
