package parsers

import (
	"fmt"
	"log"
	"strings"

	"github.com/sjsanc/encard/internal/cards"
)

func ParseMarkdown(data, deck string) ([]cards.Card, error) {
	chunks := strings.Split(string(data), "---")

	// strings.Split always returns at least 1 element
	if chunks[0] == "" {
		return nil, fmt.Errorf("deck cannot be empty: %s", deck)
	}

	var result []cards.Card

	for i, chunk := range chunks {
		lines := strings.Split(strings.TrimSpace(chunk), "\n")

		if len(lines) < 2 {
			log.Printf("[%d] Parsing error: card must have at least two lines\n", i)
			continue
		}

		front := lines[0]
		if front[0] != '#' {
			log.Printf("[%d] Parsing error: card front must start with a #\n", i)
			for _, line := range lines {
				log.Println("> " + line)
			}
			continue
		}
		front = strings.TrimPrefix(front, "# ")

		back := lines[1:]

		if back[0] == "" {
			log.Printf("[%d] Parsing error: card back must not be empty\n", i)
			continue
		}

		// Multi-choice card
		if strings.HasPrefix(back[0], "-") || strings.HasPrefix(back[0], "*") {
			choices := make([]string, 0)
			answer := -1

			for i, line := range back {
				if strings.HasPrefix(line, "-") {
					choices = append(choices, strings.TrimPrefix(line, "- "))
				} else if strings.HasPrefix(line, "*") {
					choices = append(choices, strings.TrimPrefix(line, "* "))
					answer = i
				}
			}

			card := cards.NewMultiChoice(deck, front, choices, answer)
			result = append(result, card)
			continue
		}

		// Multi-answer card
		if strings.HasPrefix(back[0], "[*]") || strings.HasPrefix(back[0], "[ ]") {
			choices := make([]string, 0)
			answers := make([]int, 0)

			for i, line := range back {
				if strings.HasPrefix(line, "[*]") {
					choices = append(choices, strings.TrimPrefix(line, "[*] "))
					answers = append(answers, i)
				} else if strings.HasPrefix(line, "[ ]") {
					choices = append(choices, strings.TrimPrefix(line, "[ ] "))
				}
			}

			card := cards.NewMultiAnswer(deck, front, choices, answers)
			result = append(result, card)
			continue
		}

		// Basic card
		card := cards.NewBasic(deck, front, strings.Join(back, "\n"))
		result = append(result, card)
	}

	return result, nil
}
