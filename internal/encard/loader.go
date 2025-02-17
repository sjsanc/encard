package encard

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/sjsanc/encard/internal/cards"
	"github.com/sjsanc/encard/internal/parsers"
)

var ENCARD_FOLDER_ROOT = "encard"

var ErrNoDefaultPath = fmt.Errorf("neither XDG_DATA_HOME nor HOME are set")

// `ResolveRootPath` returns the root path for the encard folder as a subdirectory of either $XDG_DATA_HOME or $HOME.
func ResolveRootPath() (string, error) {
	XDGDataHomePath := os.Getenv("XDG_DATA_HOME")
	homePath, _ := os.UserHomeDir()

	if XDGDataHomePath == "" && homePath == "" {
		return "", ErrNoDefaultPath
	}

	var rootPath = XDGDataHomePath

	if rootPath == "" {
		rootPath = homePath
	}

	encardPath := path.Join(rootPath, ENCARD_FOLDER_ROOT)

	_, err := os.Stat(rootPath)
	if err != nil {
		if err := os.MkdirAll(encardPath, 0755); err != nil {
			return "", err
		}
	}

	return encardPath, nil
}

// `ParseDirectory` parses a directory of files into a slice of cards.
func ParseDirectory(path string) ([]cards.Card, error) {
	var cards []cards.Card

	file, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	for _, f := range file {
		if f.IsDir() {
			continue
		}

		if filepath.Ext(f.Name()) == ".md" {
			path := filepath.Join(path, f.Name())
			deckName := strings.TrimSuffix(filepath.ToSlash(path), filepath.Ext(path))

			data, err := os.ReadFile(path)
			if err != nil {
				return nil, fmt.Errorf("error reading file: %w", err)
			}

			deck, err := parsers.ParseMarkdown(string(data), deckName)
			if err != nil {
				return nil, err
			}
			cards = append(cards, deck...)
		}
	}

	return cards, nil
}

// `LoadDeckFromPath` loads a deck of cards from a given path.
func LoadDeckFromPath(path string) ([]cards.Card, error) {
	var cards []cards.Card

	ext := filepath.Ext(path)

	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("invalid path: %w", err)
	}

	if ext == ".md" {
		deckName := strings.TrimSuffix(filepath.ToSlash(path), filepath.Ext(path))

		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("error reading file: %w", err)
		}

		deck, err := parsers.ParseMarkdown(string(data), deckName)
		if err != nil {
			return nil, err
		}
		cards = append(cards, deck...)
	} else if ext == "" {
		if !info.IsDir() {
			return nil, fmt.Errorf("invalid Deck type: %w", err)
		}
		deck, err := ParseDirectory(path)
		if err != nil {
			return nil, err
		}
		cards = append(cards, deck...)
	} else {
		return nil, fmt.Errorf("unsupported Deck type: %w", err)
	}

	if len(cards) == 0 {
		return nil, fmt.Errorf("no cards found in Deck: %s", path)
	}

	return cards, nil
}
