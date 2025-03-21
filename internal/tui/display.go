package tui

import (
	"strings"

	"github.com/sjsanc/encard/internal/defs"
	s "github.com/sjsanc/encard/internal/styles"
)

func Display(card defs.Card, faint bool) string {
	switch card.(type) {
	case *defs.Basic:
		return displayBasic(card.(*defs.Basic), faint)
	default:
		return ""
	}
}

func displayBasic(card *defs.Basic, faint bool) string {
	sb := strings.Builder{}
	sb.WriteString(s.Question.Faint(faint).Render(card.Front()) + "\n")

	if card.Flipped() {
		for _, line := range strings.Split(card.Back, "\n") {
			if strings.HasPrefix(line, "[](") {
				filepath := strings.TrimSuffix(strings.TrimPrefix(line, "[]("), ")")
				img := NewImage(filepath)
				sb.WriteString(img.Print() + "\n")
			} else {
				sb.WriteString(s.Base.Faint(faint).Render(line) + "\n")
			}
		}
	}

	return sb.String()
}
