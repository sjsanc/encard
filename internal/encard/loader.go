package encard

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobwas/glob"
	"github.com/sjsanc/encard/internal/defs"
	"github.com/sjsanc/encard/internal/parsers"
)

// `LoadCards` recursively loads globbed cards from a given root path.
func LoadCards(root string, globs []glob.Glob) ([]defs.Card, error) {
	if root == "" {
		return nil, fmt.Errorf("invalid root path")
	}

	var cards []defs.Card

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		var matched bool
		if len(globs) == 0 {
			matched = true
		} else {
			for _, g := range globs {
				if g.Match(path) {
					matched = true
				}
			}
		}
		if !matched {
			return nil
		}
		deck := strings.TrimPrefix(filepath.ToSlash(strings.TrimSuffix(path, filepath.Ext(path))), root)
		data, err := os.ReadFile(path)
		if err != nil {
			log.Printf("error reading file %s: %v\n", path, err)
			return nil
		}
		parsed, err := parsers.ParseMarkdown(string(data), deck)
		if err != nil {
			log.Printf("error parsing file %s: %v\n", path, err)
			return nil
		}
		cards = append(cards, parsed...)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %w", err)
	}

	return cards, nil
}
