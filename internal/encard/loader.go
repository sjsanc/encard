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

// Loads cards from a list of given paths.
// Absolute paths are parsed as is. Relative paths are joined with the root path.
// If a path is a directory, it will load all files in the directory and its subdirectories.
func LoadCards(paths []string, root string) ([]defs.Card, []error) {
	var errors []error
	var cards []defs.Card

	if len(paths) == 0 {
		logger.Println("no paths provided, loading all cards from", root)

		parsed, err := loadRecursive(root)
		if err != nil {
			errors = append(errors, fmt.Errorf("error loading directory %s: %v", root, err))
		}
		cards = append(cards, parsed...)
		return cards, errors
	}

	for _, path := range paths {
		target := path

		if target == "" {
			errors = append(errors, fmt.Errorf("%w: %s", ErrInvalidPath, target))
			continue
		}

		if strings.HasPrefix(target, ".") {
			target, _ = filepath.Abs(path)
		} else if !filepath.IsAbs(path) { // i.e. has prefix /
			target = filepath.Join(root, path)
		}

		logger.Println("loading cards from", target)

		info, err := os.Stat(target)
		if err != nil {
			if os.IsNotExist(err) {
				errors = append(errors, fmt.Errorf("%w: %s", ErrInvalidPath, target))
			} else {
				errors = append(errors, fmt.Errorf("error reading file %s: %v", target, err))
			}
			continue
		}

		if info.IsDir() {
			parsed, err := loadRecursive(target)
			if err != nil {
				errors = append(errors, fmt.Errorf("error loading directory %s: %v", target, err))
				continue
			}
			cards = append(cards, parsed...)

		} else {
			parsed, err := loadFromPath(target)
			if err != nil {
				errors = append(errors, fmt.Errorf("error loading file %s: %v", target, err))
				continue
			}
			cards = append(cards, parsed...)
		}
	}
	return cards, errors
}

func loadRecursive(path string) ([]defs.Card, error) {
	var cards []defs.Card

	err := filepath.WalkDir(path, func(entryPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing %s: %v", entryPath, err)
		}

		if d.IsDir() {
			return nil
		}

		parsed, err := loadFromPath(entryPath)
		if err != nil {
			return fmt.Errorf("error loading file %s: %v", entryPath, err)
		}

		cards = append(cards, parsed...)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return cards, nil
}

func loadFromPath(path string) ([]defs.Card, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", path, err)
	}

	deckname := extractDeckName(path)

	var parsed []defs.Card

	if filepath.Ext(path) == ".md" {
		parsed, err = parsers.ParseMarkdown(string(data), deckname)
		if err != nil {
			return nil, fmt.Errorf("error parsing file %s: %v", path, err)
		}
	}

	if filepath.Ext(path) == ".json" {
		parsed, err = parsers.ParseJson(string(data), deckname)
		if err != nil {
			return nil, fmt.Errorf("error parsing file %s: %v", path, err)
		}
	}

	logger.Printf("loaded %d cards from %s (%s)", len(parsed), deckname, path)

	return parsed, nil
}

func extractDeckName(path string) string {
	return strings.TrimSuffix(filepath.Base(filepath.Dir(path)), filepath.Ext(path))
}
