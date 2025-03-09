package defs

import (
	"strings"

	"github.com/sjsanc/encard/internal/styles"
)

type Cloze struct {
	Base
	text     []string // The text of the cloze card, split into segments
	input    map[int]string
	keys     []int
	selected int
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
		text:     text,
		input:    input,
		keys:     keys,
		selected: keys[0],
	}
}

func (c *Cloze) Update(key string) {
	switch key {
	case "backspace":
		val := c.input[c.selected]
		if len(val) > 0 {
			c.input[c.selected] = val[:len(val)-1]
		}
	case "enter":
		c.Flip()
	case "tab", "right":
		for i, k := range c.keys {
			if k == c.selected {
				c.selected = c.keys[(i+1)%len(c.keys)]
				break
			}
		}
	case "left":
		for i, k := range c.keys {
			if k == c.selected {
				c.selected = c.keys[(i-1+len(c.keys))%len(c.keys)]
				break
			}
		}
	default:
		c.input[c.selected] += key
	}
}

func (c *Cloze) Render(faint bool) string {
	sb := strings.Builder{}
	sb.WriteString(styles.Question.Faint(faint).Render(c.front) + "\n")

	if c.flipped {
		for i, segment := range c.text {
			if val, ok := c.input[i]; ok {
				answer := strings.TrimPrefix(segment, "{{")
				answer = strings.TrimSuffix(answer, "}}")
				if val == answer {
					sb.WriteString(styles.Correct.Render(val) + " ")
				} else {
					sb.WriteString(styles.Incorrect.Strikethrough(true).Render(val) + " ")
					sb.WriteString(styles.IncorrectUnselected.Render(answer) + " ")
				}
			} else {
				sb.WriteString(segment + " ")
			}
		}
	} else {
		for i, segment := range c.text {
			if val, ok := c.input[i]; ok {
				if i == c.selected {
					sb.WriteString(styles.Selected.Render("_" + val + "_ "))
				} else {
					sb.WriteString("_" + val + "_ ")
				}
			} else {
				sb.WriteString(segment + " ")
			}
		}
	}

	return sb.String()
}
