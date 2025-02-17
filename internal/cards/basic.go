package cards

import (
	"strings"

	"github.com/sjsanc/encard/internal/styles"
)

type Basic struct {
	CardBase
	back string
}

func NewBasic(deck string, front string, back string) *Basic {
	return &Basic{
		CardBase: CardBase{
			deck:  deck,
			front: front,
		},
		back: back,
	}
}

func (c *Basic) Render() string {
	sb := strings.Builder{}
	sb.WriteString(styles.Question.Render(c.front) + "\n")
	if c.flipped {
		sb.WriteString(c.back)
	}
	return sb.String()
}
