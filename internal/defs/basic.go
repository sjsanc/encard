package defs

type Basic struct {
	Base
	Deck  string
	Front string
	Back  string
}

func NewBasic(deck string, front string, back string) *Basic {
	return &Basic{
		Deck:  deck,
		Front: front,
		Back:  back,
	}
}

func (c *Basic) Update(key string) {
	switch key {
	case "enter":
		c.Flip()
	}
}
