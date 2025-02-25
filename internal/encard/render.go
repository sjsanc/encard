package encard

import (
	"github.com/sjsanc/encard/internal/defs"
)

func RenderCard(card defs.Card, flipped bool) string {
	switch c := card.(type) {
	case *defs.Basic:
		return renderBasicCard(c, flipped)
	case *defs.Input:
		return renderInputCard(c, flipped)
	case *defs.MultiChoice:
		return renderMultiChoiceCard(c, flipped)
	case *defs.MultiAnswer:
		return renderMultiAnswerCard(c, flipped)
	}
	return ""
}

func renderBasicCard(card *defs.Basic, flipped bool) string {
	// sb := strings.Builder{}

	// sb.WriteString(styles.Question.Render(card.front) + "\n")

	// if flipped {
	// 	sb.WriteString(c.back)
	// }

	// result := lipgloss.NewStyle().
	// 	Padding(0, 1).
	// 	Border(lipgloss.NormalBorder(), false, false, false, true).
	// 	BorderForeground(lipgloss.Color("#542b7c")).
	// 	Render(sb.String())

	// return result
	return ""
}

func renderInputCard(card *defs.Input, flipped bool) string {
	// sb := strings.Builder{}

	// sb.WriteString(styles.Question.Render(c.front) + "\n")

	// distance := levenshtein.ComputeDistance(c.input, c.answer)
	// isCorrect := threshold(len(c.answer), distance)

	// if c.flipped {
	// 	if isCorrect {
	// 		sb.WriteString(styles.Correct.Render(c.input) + "\n")
	// 	} else {
	// 		sb.WriteString(styles.Incorrect.Render(c.input) + "\n")
	// 	}
	// 	sb.WriteString(styles.Selected.Render(c.answer) + "\n")
	// }

	// return sb.String()
	return ""
}

func renderMultiChoiceCard(card *defs.MultiChoice, flipped bool) string {
	// sb := strings.Builder{}
	// sb.WriteString(styles.Question.Render(c.front) + "\n")

	// for i, choice := range c.choices {
	// 	if c.flipped {
	// 		// Selected + Correct
	// 		if c.current == i && c.answer == i {
	// 			sb.WriteString(styles.Correct.Render("* "+choice+" (correct!)") + "\n")
	// 		}

	// 		// Selected + Incorrect
	// 		if c.current == i && c.answer != i {
	// 			sb.WriteString(styles.Incorrect.Render("* "+choice+" (incorrect!)") + "\n")
	// 		}

	// 		// Not Selected + Correct
	// 		if c.current != i && c.answer == i {
	// 			sb.WriteString(styles.IncorrectUnselected.Render("- "+choice+" (answer)") + "\n")
	// 		}

	// 		// Not Selected + Incorrect
	// 		if c.current != i && c.answer != i {
	// 			sb.WriteString("- " + choice + "\n")
	// 		}
	// 	} else {
	// 		if c.current == i {
	// 			sb.WriteString(styles.Selected.Render("* "+choice) + "\n")
	// 		} else {
	// 			sb.WriteString("- " + choice + "\n")
	// 		}
	// 	}
	// }

	// return sb.String()
	return ""
}

func renderMultiAnswerCard(c *defs.MultiAnswer, flipped bool) string {
	// sb := strings.Builder{}
	// sb.WriteString(styles.Question.Render(c.front) + "\n")

	// for i, choice := range c.choices {
	// 	if c.flipped {
	// 		// Selected + Correct
	// 		if slices.Contains(c.selected, i) && slices.Contains(c.answers, i) {
	// 			sb.WriteString(styles.Correct.Render("[x] "+choice+" (correct!)") + "\n")
	// 		}

	// 		// Selected + Incorrect
	// 		if slices.Contains(c.selected, i) && !slices.Contains(c.answers, i) {
	// 			sb.WriteString(styles.Incorrect.Render("[x] "+choice+" (incorrect!)") + "\n")
	// 		}

	// 		// Not Selected + Correct
	// 		if !slices.Contains(c.selected, i) && slices.Contains(c.answers, i) {
	// 			sb.WriteString(styles.IncorrectUnselected.Render("[ ] "+choice+" (answer)") + "\n")
	// 		}

	// 		// Not Selected + Incorrect
	// 		if !slices.Contains(c.selected, i) && !slices.Contains(c.answers, i) {
	// 			sb.WriteString("[ ] " + choice + "\n")
	// 		}
	// 	} else {
	// 		if slices.Contains(c.selected, i) && c.current == i {
	// 			sb.WriteString(styles.Selected.Render("[x] "+choice) + "\n")
	// 		} else if slices.Contains(c.selected, i) {
	// 			sb.WriteString("[x] " + choice + "\n")
	// 		} else if c.current == i {
	// 			sb.WriteString(styles.Selected.Render("[ ] "+choice) + "\n")
	// 		} else {
	// 			sb.WriteString("[ ] " + choice + "\n")
	// 		}
	// 	}
	// }

	// return sb.String()
	return ""
}
