package tui

import (
	"strings"

	"github.com/sjsanc/encard/internal/defs"
	s "github.com/sjsanc/encard/internal/styles"
)

func displayCard(c defs.Card, f bool) string {
	switch card := c.(type) {
	case *defs.Basic:
		return displayBasic(card, f)
	case *defs.Cloze:
		return displayCloze(card, f)
	case *defs.Input:
		return displayInput(card, f)
	case *defs.MultiAnswer:
		return displayMultiAnswer(card, f)
	case *defs.MultiChoice:
		return displayMultiChoice(card, f)
	default:
		return ""
	}
}

func displayBasic(c *defs.Basic, f bool) string {
	sb := strings.Builder{}
	sb.WriteString(s.Question.Faint(f).Render(c.Front()) + "\n")

	if c.Flipped() {
		for _, line := range strings.Split(c.Back, "\n") {
			if strings.HasPrefix(line, "[](") {
				filepath := strings.TrimSuffix(strings.TrimPrefix(line, "[]("), ")")
				img := NewImage(filepath)
				sb.WriteString(img.Print() + "\n")
			} else {
				sb.WriteString(s.Base.Faint(f).Render(line) + "\n")
			}
		}
	}

	return sb.String()
}

func displayCloze(c *defs.Cloze, f bool) string {
	sb := strings.Builder{}
	sb.WriteString(s.Question.Faint(f).Render(c.Front()) + "\n")

	if c.Flipped() {
		for i, seg := range c.Text {
			if val, ok := c.Input[i]; ok {
				ans := strings.TrimPrefix(seg, "{{")
				ans = strings.TrimSuffix(ans, "}}")
				if val == ans {
					sb.WriteString(s.Correct.Render(val) + " ")
				} else {
					sb.WriteString(s.Incorrect.Strikethrough(true).Render(val) + " ")
					sb.WriteString(s.IncorrectUnselected.Render(ans) + " ")
				}
			} else {
				sb.WriteString(seg + " ")
			}
		}
	} else {
		for i, seg := range c.Text {
			if val, ok := c.Input[i]; ok {
				if i == c.Selected {
					sb.WriteString(s.Selected.Render("_" + val + "_ "))
				} else {
					sb.WriteString("_" + val + "_ ")
				}
			} else {
				sb.WriteString(seg + " ")
			}
		}
	}

	return sb.String()
}

func displayInput(c *defs.Input, f bool) string {
	sb := strings.Builder{}
	sb.WriteString(s.Question.Faint(f).Render(c.Front()) + "\n")

	if c.Flipped() {
		sb.WriteString(s.Base.Render(c.Input) + "\n")
		if c.Matched {
			sb.WriteString(s.Correct.Render(c.Answer) + "\n")
		} else {
			sb.WriteString(s.Incorrect.Render(c.Answer) + "\n")
		}
	} else {
		sb.WriteString(s.Base.Render(c.Input) + "\n")
	}

	return sb.String()
}

func displayMultiAnswer(c *defs.MultiAnswer, f bool) string {
	sb := strings.Builder{}
	sb.WriteString(s.Question.Faint(f).Render(c.Front()) + "\n")

	for i, choice := range c.Answers {
		selected := choice.Selected
		correct := choice.Correct

		if c.Flipped() {
			if selected && correct {
				sb.WriteString(s.Correct.Faint(f).Render("[x] "+choice.Text+" (correct!)") + "\n")
			} else if selected && !correct {
				sb.WriteString(s.Incorrect.Faint(f).Render("[x] "+choice.Text+" (incorrect!)") + "\n")
			} else if !selected && correct {
				sb.WriteString(s.IncorrectUnselected.Faint(f).Render("[ ] "+choice.Text+" (answer)") + "\n")
			} else {
				sb.WriteString(s.Base.Faint(f).Render("[ ] "+choice.Text) + "\n")
			}
		} else {
			if c.Current == i {
				if selected {
					sb.WriteString(s.Selected.Render("[x] "+choice.Text) + "\n")
				} else {
					sb.WriteString(s.Selected.Render("[ ] "+choice.Text) + "\n")
				}
			} else {
				if selected {
					sb.WriteString(s.Base.Faint(f).Render("[x] "+choice.Text) + "\n")
				} else {
					sb.WriteString(s.Base.Faint(f).Render("[ ] "+choice.Text) + "\n")
				}
			}
		}
	}

	return sb.String()
}

func displayMultiChoice(c *defs.MultiChoice, f bool) string {
	sb := strings.Builder{}
	sb.WriteString(s.Question.Faint(f).Render(c.Front()) + "\n")

	for i, choice := range c.Choices {
		if c.Flipped() {
			if c.Current == i && choice.Correct {
				sb.WriteString(s.Correct.Faint(f).Render("* "+choice.Text+" (correct!)") + "\n")
			} else if c.Current == i && !choice.Correct {
				sb.WriteString(s.Incorrect.Faint(f).Render("* "+choice.Text+" (incorrect!)") + "\n")
			} else if c.Current != i && choice.Correct {
				sb.WriteString(s.IncorrectUnselected.Faint(f).Render("- "+choice.Text+" (answer)") + "\n")
			} else {
				sb.WriteString("- " + choice.Text + "\n")
			}
		} else {
			if c.Current == i {
				sb.WriteString(s.Selected.Faint(f).Render("* "+choice.Text) + "\n")
			} else {
				sb.WriteString("- " + choice.Text + "\n")
			}
		}
	}

	return sb.String()
}
