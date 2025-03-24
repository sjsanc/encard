package defs

import (
	"strings"
)

type Cloze struct {
	Base
	Text     []string // The text of the cloze card, split into segments
	Input    map[int]string
	Keys     []int
	Selected int
}

func NewCloze(deck string, front string, text []string) *Cloze {
	input := make(map[int]string)
	for i := range text {
		if strings.Contains(text[i], "{{") {
			input[i] = ""
		}
	}

	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}

	return &Cloze{
		Base: Base{
			variant: "cloze",
			deck:    deck,
			front:   front,
		},
		Text:     text,
		Input:    input,
		Keys:     keys,
		Selected: keys[0],
	}
}

func (c *Cloze) Update(key string) {
	switch key {
	case "backspace":
		val := c.Input[c.Selected]
		if len(val) > 0 {
			c.Input[c.Selected] = val[:len(val)-1]
		}
	case "enter":
		c.Flip()
	case "tab", "right":
		for i, k := range c.Keys {
			if k == c.Selected {
				c.Selected = c.Keys[(i+1)%len(c.Keys)]
				break
			}
		}
	case "left":
		for i, k := range c.Keys {
			if k == c.Selected {
				c.Selected = c.Keys[(i-1+len(c.Keys))%len(c.Keys)]
				break
			}
		}
	default:
		c.Input[c.Selected] += key
	}
}
