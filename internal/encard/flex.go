package encard

import (
	"strings"

	lg "github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

var ns = lg.NewStyle()

type Element struct {
	Text     string
	BG       lg.Color
	FG       lg.Color
	PaddingL int
	PaddingR int
	Grow     bool
}

func NewElement(text string, bg, fg string, Grow bool) Element {
	return Element{
		Text:     text,
		BG:       lg.Color(bg),
		FG:       lg.Color(fg),
		PaddingL: 1,
		PaddingR: 1,
		Grow:     Grow,
	}
}

func (e *Element) Width() int {
	return lg.Width(e.Text) + e.PaddingL + e.PaddingR
}

type Flex struct {
	Children []Element
	Wrap     bool
}

func calculateWidths(children []Element, maxWidth int) []int {
	widths := make([]int, len(children))
	var growIndices []int
	fixedWidthSum := 0

	// First pass: record fixed widths and grow indices.
	for i, c := range children {
		if c.Grow {
			growIndices = append(growIndices, i)
		} else {
			widths[i] = c.Width()
			fixedWidthSum += widths[i]
		}
	}

	// Compute remaining space and distribute it.
	growWidth := maxWidth - fixedWidthSum
	growCount := len(growIndices)
	if growCount > 0 {
		unitGrowWidth := growWidth / growCount
		extraPixels := growWidth % growCount
		for _, i := range growIndices {
			widths[i] = unitGrowWidth
			if extraPixels > 0 {
				widths[i]++
				extraPixels--
			}
		}
	}

	return widths
}

func (f *Flex) Render(width int) string {
	widths := calculateWidths(f.Children, width)

	maxHeight := 1

	lines := make(map[int]*strings.Builder)

	// Determine max height required
	wrappedTexts := make([][]string, len(f.Children))
	for i, c := range f.Children {
		contentWidth := widths[i] - c.PaddingL - c.PaddingR
		if contentWidth < 0 {
			contentWidth = 0
		}

		wrappedTexts[i] = strings.Split(wordwrap.String(c.Text, contentWidth), "\n")

		if len(wrappedTexts[i]) > maxHeight {
			maxHeight = len(wrappedTexts[i])
		}
	}

	// Render each element
	for i, c := range f.Children {
		for j := 0; j < maxHeight; j++ {
			if _, exists := lines[j]; !exists {
				lines[j] = &strings.Builder{}
			}

			text := ""
			if j < len(wrappedTexts[i]) {
				text = wrappedTexts[i][j]
			} else {
				text = strings.Repeat(" ", widths[i]-c.PaddingL-c.PaddingR)
			}

			lines[j].WriteString(ns.
				Foreground(c.FG).
				PaddingLeft(c.PaddingL).
				PaddingRight(c.PaddingR).
				Background(c.BG).
				Width(widths[i]).
				Render(text))
		}
	}

	// Build final output
	var sb strings.Builder
	for j := 0; j < maxHeight; j++ { // Ensure proper order
		sb.WriteString(lines[j].String())
		sb.WriteRune('\n')
	}

	return ns.Render(sb.String())
}
