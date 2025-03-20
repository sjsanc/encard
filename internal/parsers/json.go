package parsers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sjsanc/encard/internal/defs"
)

type CardData struct {
	Front   string          `json:"front"`
	Back    string          `json:"back,omitempty"`
	Type    string          `json:"type,omitempty"`
	Answers map[string]bool `json:"answers,omitempty"`
	Choices map[string]bool `json:"choices,omitempty"`
}

func ParseJson(data string, deck string) ([]defs.Card, error) {
	var results []defs.Card

	var cards []CardData
	if err := json.Unmarshal([]byte(data), &cards); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	for _, obj := range cards {
		var card defs.Card
		switch obj.Type {
		case "cloze":
			text := strings.Fields(obj.Back)
			card = defs.NewCloze(deck, obj.Front, text)
		case "input":
			card = defs.NewInput(deck, obj.Front, obj.Back)
		case "multianswer":
			card = defs.NewMultiAnswer(deck, obj.Front, obj.Answers)
		case "multichoice":
			card = defs.NewMultiChoice(deck, obj.Front, obj.Choices)
		default:
			card = defs.NewBasic(deck, obj.Front, obj.Back, "")
		}

		results = append(results, card)
	}

	return results, nil
}
