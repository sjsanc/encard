package encard

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/sjsanc/encard/internal/defs"
	"github.com/sjsanc/encard/internal/log"
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
		log.Info("no paths provided, loading all cards from %s", root)

		parsed, err := loadRecursive(root, root)
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

		log.Info("loading cards from %s", target)

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
			parsed, err := loadRecursive(target, target)
			if err != nil {
				errors = append(errors, fmt.Errorf("error loading directory %s: %v", target, err))
				continue
			}
			cards = append(cards, parsed...)

		} else {
			parsed, err := loadFromPath(target, filepath.Dir(target))
			if err != nil {
				errors = append(errors, fmt.Errorf("error loading file %s: %v", target, err))
				continue
			}
			cards = append(cards, parsed...)
		}
	}
	return cards, errors
}

func loadRecursive(path string, root string) ([]defs.Card, error) {
	var cards []defs.Card

	err := filepath.WalkDir(path, func(entryPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing %s: %v", entryPath, err)
		}

		if d.IsDir() {
			return nil
		}

		parsed, err := loadFromPath(entryPath, root)
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

func loadFromPath(path string, root string) ([]defs.Card, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", path, err)
	}

	deckname := extractDeckName(path, root)

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

	log.Info("loaded %d cards from %s", len(parsed), deckname)

	return parsed, nil
}

func extractDeckName(path string, root string) string {
	// Get the directory containing the file
	dir := filepath.Dir(path)

	// Try to get the relative path from root
	relPath, err := filepath.Rel(root, dir)
	if err != nil || relPath == "." {
		// If we can't get a relative path, just use the directory basename
		// This handles cases where the file is directly in the root
		baseName := filepath.Base(dir)
		return strings.TrimSuffix(baseName, filepath.Ext(baseName))
	}

	// Normalize path separators to forward slashes for cross-platform consistency
	deckName := strings.ReplaceAll(relPath, string(filepath.Separator), "/")

	// Remove common file extensions from the final path component if present
	// This handles cases like "spanish.md" -> "spanish"
	if ext := filepath.Ext(deckName); ext == ".md" || ext == ".json" {
		deckName = strings.TrimSuffix(deckName, ext)
	}

	return deckName
}
