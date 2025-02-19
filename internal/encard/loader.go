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
func LoadCards(rootPath string, globs []glob.Glob) (map[string][]cards.Card, error) {
	if rootPath == "" {
		return nil, fmt.Errorf("invalid root path")
	}

	var wg sync.WaitGroup
	m := sync.Map{}
	fileChan := make(chan string, 100)
	const workerCount = 25

	worker := func() {
		for path := range fileChan {
			ext := filepath.Ext(path)
			deck := strings.TrimPrefix(strings.TrimPrefix(filepath.ToSlash(strings.TrimSuffix(path, ext)), rootPath), "/")

			data, err := os.ReadFile(path)
			if err != nil {
				fmt.Printf("error reading file: %v\n", err)
				continue
			}

			var cards []cards.Card
			if ext == ".md" {
				c, _ := parsers.ParseMarkdown(string(data))
				cards = c
			}
			if len(cards) > 0 {
				m.Store(deck, cards)
			}
		}
		wg.Done()
	}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker()
	}

	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking directory: %w", err)
		}
		if !d.IsDir() {
			fileChan <- path
		}
		return nil
	})
	close(fileChan)
	wg.Wait()

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %w", err)
	}

	result := map[string][]cards.Card{}
	m.Range(func(key, value interface{}) bool {
		result[key.(string)] = value.([]cards.Card)
		return true
	})

	return result, nil
}
