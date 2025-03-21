package defs

// When parsing markdown, we want to go through each line.
// When parsing JSON, we only need to ensure the deserialised object is correct.
// When rendering Card data, we need to keep track of content width and image heights.
// In effect, rendering needs to be deoupled from the data structure.
// Especially important for implementing different rendering methods, such as a Web UI.

type Basic struct {
	Base
	Back string
}

func NewBasic(deck string, front string, back string) *Basic {
	return &Basic{
		Base: Base{
			variant: "basic",
			deck:    deck,
			front:   front,
		},
		Back: back,
	}
}

func (c *Basic) Update(key string) {
	switch key {
	case "enter":
		c.Flip()
	}
}
