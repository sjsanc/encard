package encard

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gobwas/glob"
	"github.com/sjsanc/encard/internal/cards"
	"github.com/sjsanc/encard/internal/parsers"
)

// `LoadCards` recursively loads globbed cards from a given root path.
func LoadCards(rootPath string, globs []glob.Glob) ([]cards.Card, error) {
	if rootPath == "" {
		return nil, fmt.Errorf("invalid root path")
	}

	var wg sync.WaitGroup
	resultChan := make(chan []cards.Card)

	processFile := func(path string) {
		defer wg.Done()
		ext := filepath.Ext(path)
		deckName := strings.TrimSuffix(filepath.ToSlash(path), ext)
		deckName = strings.TrimPrefix(deckName, rootPath)

		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", path, err)
			return
		}

		var cards []cards.Card
		if ext == ".md" {
			deck, _ := parsers.ParseMarkdown(string(data), deckName)
			cards = append(cards, deck...)
		}

		resultChan <- cards
	}

	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking directory: %w", err)
		}

		if !d.IsDir() {
			wg.Add(1)
			go processFile(path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %w", err)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var allCards []cards.Card
	for cards := range resultChan {
		allCards = append(allCards, cards...)
	}

	return allCards, nil
}
