package encard

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/sjsanc/encard/internal/defs"
	"github.com/sjsanc/encard/internal/parsers"
)

var ErrInvalidPath = fmt.Errorf("file does not exist")

// LoadCards handles loading cards from different path types
func LoadCards(paths []string, cfg *Config) ([]defs.Card, []error) {
	var errors []error
	var cards []defs.Card

	for _, path := range paths {
		targetPath := path
		if !filepath.IsAbs(path) {
			targetPath = filepath.Join(cfg.CardsDir, path)
		}

		info, err := os.Stat(targetPath)
		if err != nil {
			if os.IsNotExist(err) {
				errors = append(errors, fmt.Errorf("%w: %s", ErrInvalidPath, targetPath))
			} else {
				errors = append(errors, fmt.Errorf("error reading file %s: %v", targetPath, err))
			}
			continue
		}

		if info.IsDir() {
			filepath.WalkDir(targetPath, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					errors = append(errors, fmt.Errorf("error accessing %s: %v", path, err))
					return nil
				}
				if d.IsDir() {
					return nil
				}
				parsed, err := loadFromPath(path)
				if err != nil {
					errors = append(errors, fmt.Errorf("error loading file %s: %v", path, err))
					return nil
				}
				cards = append(cards, parsed...)
				return nil
			})
		} else {
			parsed, err := loadFromPath(targetPath)
			if err != nil {
				errors = append(errors, fmt.Errorf("error loading file %s: %v", targetPath, err))
				continue
			}
			cards = append(cards, parsed...)
		}
	}

	return cards, errors
}

func loadFromPath(path string) ([]defs.Card, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", path, err)
	}

	deckname := extractDeckName(path)
	parsed, err := parsers.ParseMarkdown(string(data), deckname)
	if err != nil {
		return nil, fmt.Errorf("error parsing file %s: %v", path, err)
	}

	return parsed, nil
}

func extractDeckName(path string) string {
	return strings.TrimSuffix(filepath.Base(filepath.Dir(path)), filepath.Ext(path))
}
