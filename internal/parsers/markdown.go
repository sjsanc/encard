package parsers

import (
	"log"
	"strings"

	"github.com/sjsanc/encard/internal/defs"
)

func ParseMarkdown(data string, deck string) ([]defs.Card, error) {
	chunks := strings.Split(string(data), "---")

	// strings.Split always returns at least 1 element
	if chunks[0] == "" {
		log.Printf("Parsing error: no cards found in %s\n", deck)
		return nil, nil
	}

	var result []defs.Card

	for i, chunk := range chunks {
		if chunk == "" {
			continue // ignore the empty chunks in the case of single card files
		}

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

		if strings.HasSuffix(front, "{}") {
			cardA := defs.NewBasic(deck, strings.Replace(front, "{}", back[0], -1), back[1])
			cardB := defs.NewBasic(deck, strings.Replace(front, "{}", back[1], -1), back[0])
			result = append(result, cardA, cardB)
			continue
		}

		// Multi-choice card
		if strings.HasPrefix(back[0], "-") || strings.HasPrefix(back[0], "*") {
			choices := make(map[string]bool)

			for _, line := range back {
				if strings.HasPrefix(line, "-") {
					choices[strings.TrimPrefix(line, "- ")] = false
				} else if strings.HasPrefix(line, "*") {
					choices[strings.TrimPrefix(line, "* ")] = true
				}
			}

			card := defs.NewMultiChoice(deck, front, choices)
			result = append(result, card)
			continue
		}

		// Multi-answer card
		if strings.HasPrefix(back[0], "[*]") || strings.HasPrefix(back[0], "[ ]") {
			choices := make(map[string]bool)

			for _, line := range back {
				if strings.HasPrefix(line, "[*]") {
					choices[strings.TrimPrefix(line, "[*] ")] = true
				} else if strings.HasPrefix(line, "[ ]") {
					choices[strings.TrimPrefix(line, "[ ] ")] = false
				}
			}

			card := defs.NewMultiAnswer(deck, front, choices)
			result = append(result, card)
			continue
		}

		// Input card
		if strings.HasPrefix(back[0], ">") {
			card := defs.NewInput(deck, front, strings.Join(back, "\n"))
			result = append(result, card)
			continue
		}

		// Cloze card
		if strings.Contains(back[0], "{{") {
			text := strings.Fields(back[0])
			card := defs.NewCloze(deck, front, text)
			result = append(result, card)
			continue
		}

		// for _, line := range back {
		// 	if strings.HasPrefix("[](", line) {

		// 	}
		// }

		// Basic card
		card := defs.NewBasic(deck, front, strings.Join(back, "\n"))
		result = append(result, card)
	}

	return result, nil
}
