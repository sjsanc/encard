package flex

// import (
// 	"strings"

// 	lg "github.com/charmbracelet/lipgloss"
// 	"github.com/muesli/reflow/wordwrap"
// )

// var ns = lg.NewStyle()

// type Text struct {
// 	Text string
// 	BG   lg.Color
// 	FG   lg.Color
// 	Bold bool
// }

// type Element struct {
// 	Nodes    []Text
// 	BG       lg.Color
// 	FG       lg.Color
// 	PaddingL int
// 	PaddingR int
// 	Grow     bool
// }

// func NewElement(nodes []Text, bg, fg string, Grow bool) Element {
// 	return Element{
// 		Nodes:    nodes,
// 		BG:       lg.Color(bg),
// 		FG:       lg.Color(fg),
// 		PaddingL: 1,
// 		PaddingR: 1,
// 		Grow:     Grow,
// 	}
// }

// func (e *Element) Width() int {
// 	width := 0
// 	for _, t := range e.Nodes {
// 		w := lg.Width(t.Text)
// 		if w > width {
// 			width = w
// 		}
// 	}
// 	return width + e.PaddingL + e.PaddingR
// }

// type Flex struct {
// 	Children []Element
// 	Wrap     bool
// }

// func calculateWidths(children []Element, maxWidth int) []int {
// 	widths := make([]int, len(children))
// 	var growIndices []int
// 	fixedWidthSum := 0

// 	// First pass: record fixed widths and grow indices.
// 	for i, c := range children {
// 		if c.Grow {
// 			growIndices = append(growIndices, i)
// 		} else {
// 			widths[i] = c.Width()
// 			fixedWidthSum += widths[i]
// 		}
// 	}

// 	// Compute remaining space and distribute it.
// 	growWidth := maxWidth - fixedWidthSum
// 	growCount := len(growIndices)
// 	if growCount > 0 {
// 		unitGrowWidth := growWidth / growCount
// 		extraPixels := growWidth % growCount
// 		for _, i := range growIndices {
// 			widths[i] = unitGrowWidth
// 			if extraPixels > 0 {
// 				widths[i]++
// 				extraPixels--
// 			}
// 		}
// 	}

// 	return widths
// }

// func (f *Flex) Render(width int) string {
// 	widths := calculateWidths(f.Children, width)

// 	maxHeight := 1
// 	wrappedTexts := make([][]string, len(f.Children))

// 	// Determine max height required
// 	for i, c := range f.Children {
// 		if len(c.Nodes) == 0 {
// 			continue
// 		}

// 		contentWidth := widths[i] - c.PaddingL - c.PaddingR
// 		if contentWidth < 1 {
// 			contentWidth = 1
// 		}

// 		var wrappedLines []string

// 		for _, t := range c.Nodes {
// 			wrapped := strings.Split(wordwrap.String(t.Text, contentWidth), "\n")
// 			for _, line := range wrapped {
// 				wrappedLines = append(wrappedLines, ns.Foreground(t.FG).Background(t.BG).Render(line))
// 			}
// 		}

// 		wrappedTexts[i] = wrappedLines
// 		if len(wrappedLines) > maxHeight {
// 			maxHeight = len(wrappedLines)
// 		}
// 	}

// 	// Ensure all elements have the same height
// 	for i := range f.Children {
// 		if len(wrappedTexts[i]) < maxHeight {
// 			for j := len(wrappedTexts[i]); j < maxHeight; j++ {
// 				wrappedTexts[i] = append(wrappedTexts[i], "")
// 			}
// 		}
// 	}

// 	lines := make(map[int]*strings.Builder)

// 	// Render each element
// 	for j := 0; j < maxHeight; j++ {
// 		lines[j] = &strings.Builder{}

// 		for i, c := range f.Children {
// 			text := wrappedTexts[i][j]

// 			// Apply element-level styling
// 			lines[j].WriteString(ns.
// 				Foreground(c.FG).
// 				Background(c.BG).
// 				PaddingLeft(c.PaddingL).
// 				PaddingRight(c.PaddingR).
// 				Width(widths[i]).
// 				Render(text))
// 		}
// 	}

// 	// Build final output
// 	var sb strings.Builder
// 	for j := 0; j < maxHeight; j++ { // Ensure proper order
// 		sb.WriteString(lines[j].String())
// 		// sb.WriteRune('\n')
// 	}

// 	return ns.Render(sb.String())
// }

// package encard

// import (
// 	"strings"

// 	lg "github.com/charmbracelet/lipgloss"
// 	"github.com/muesli/reflow/wordwrap"
// )

// type FlexDir int

// const (
// 	FlexCol FlexDir = iota
// 	FlexRow
// )

// type Node struct {
// 	text     string
// 	style    lg.Style
// 	children []*Node
// 	grow     bool
// 	width    int // Fixed width
// 	flexDir  FlexDir
// }

// func NewNode(t string) *Node {
// 	return &Node{
// 		text:     t,
// 		children: make([]*Node, 0),
// 	}
// }

// func Row() *Node {
// 	return NewNode("").Flex(FlexRow)
// }
// func Col() *Node {
// 	return NewNode("").Flex(FlexCol)
// }

// func (n *Node) AddChild(c *Node) *Node {
// 	n.children = append(n.children, c)
// 	return n
// }
// func (n *Node) Flex(d FlexDir) *Node {
// 	n.flexDir = d
// 	return n
// }
// func (n *Node) Style(s lg.Style) *Node {
// 	n.style = s
// 	return n
// }
// func (n *Node) Grow() *Node {
// 	n.grow = true
// 	return n
// }
// func (n *Node) Width(w int) *Node {
// 	n.width = w
// 	return n
// }

// func (n *Node) Render(totalW int) string {
// 	// Render leaf node as text
// 	if len(n.children) == 0 {
// 		return n.style.Width(totalW).Render(wordwrap.String(n.text, totalW))
// 	}

// 	widths := make([]int, len(n.children))

// 	if n.flexDir == FlexCol {
// 		for i := range n.children {
// 			widths[i] = totalW
// 		}
// 	} else {
// 		fixedWidthSum := 0
// 		growIndices := make([]int, 0)

// 		for i, c := range n.children {
// 			if c.grow {
// 				growIndices = append(growIndices, i)
// 				continue
// 			}
// 			var w int
// 			if c.width > 0 {
// 				w = c.width // Use fixed width
// 			} else {
// 				w = lg.Width(c.text) + c.style.GetPaddingLeft() + c.style.GetPaddingRight()
// 			}
// 			widths[i] = w
// 			fixedWidthSum += w
// 		}

// 		growCount := len(growIndices)
// 		if growCount > 0 {
// 			growWidth := totalW - fixedWidthSum
// 			growBy := growWidth / growCount
// 			growExtra := growWidth % growCount

// 			for _, i := range growIndices {
// 				widths[i] = growBy
// 				if growExtra > 0 {
// 					widths[i]++
// 					growExtra--
// 				}
// 			}
// 		}
// 	}

// 	if n.flexDir == FlexCol {
// 		rows := make([]string, len(n.children))
// 		maxHeight := 1
// 		for i, c := range n.children {
// 			text := c.Render(widths[i])
// 			lines := strings.Split(text, "\n")
// 			if len(lines) > maxHeight {
// 				maxHeight = len(lines)
// 			}
// 			rows[i] = c.Render(widths[i])
// 		}
// 		for i := len(rows); i < maxHeight; i++ {
// 			rows = append(rows, "")
// 		}
// 		return lg.JoinVertical(lg.Top, rows...)
// 	}

// 	cols := make([]string, len(n.children))
// 	maxHeight := 1

// 	for i, c := range n.children {
// 		lines := strings.Split(c.Render(widths[i]), "\n")
// 		if len(lines) > maxHeight {
// 			maxHeight = len(lines)
// 		}
// 		cols[i] = c.Render(widths[i])
// 	}
// 	for i := len(cols); i < maxHeight; i++ {
// 		cols = append(cols, "")
// 	}

// 	return lg.JoinHorizontal(lg.Top, cols...)
// }

// var s1 = ns.Background(lg.Color("#EC4899")).Padding(0, 1)
// var s2 = ns.Background(lg.Color("#202226")).Padding(0, 1)
// var s3 = ns.Background(lg.Color("#A855F7")).Padding(0, 1)

// func (m *Model) flex2() string {

// 	main := Col()

// 	bar := Row()
// 	bar.AddChild(NewNode("A Very Long Text That Will Wrap").Style(s1))
// 	bar.AddChild(NewNode("").Grow().Style(s2))
// 	bar.AddChild(NewNode("Deck 1/2").Style(s3))

// 	body := Row()
// 	body.AddChild(NewNode("").Style(ns.Background(lg.Color("#EC4899"))).Grow())
// 	body.AddChild(NewNode("Card"))
// 	body.AddChild(NewNode("").Grow())

// 	main.AddChild(bar)
// 	main.AddChild(body)

// 	return main.Render(m.Width)
// }
