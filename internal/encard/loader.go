package encard

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
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

// `ParseMarkdownFile` parses a markdown file into a slice of cards.
func ParseMarkdownFile(path string) ([]Card, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	chunks := strings.Split(string(data), "---")

	// strings.Split always returns at least 1 element
	if chunks[0] == "" {
		return nil, fmt.Errorf("file cannot be empty: %s", path)
	}

	cards := make([]Card, 0)

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

		if strings.HasPrefix(back[0], "-") || strings.HasPrefix(back[0], "*") {
			ext := filepath.Ext(path)
			deckName := strings.TrimSuffix(filepath.ToSlash(path), ext)

			c := &MultipleChoiceCard{
				deck:    deckName,
				Front:   front,
				Choices: make([]string, 0),
			}

			choices := make([]string, 0)
			for i, line := range back {
				if strings.HasPrefix(line, "-") {
					choices = append(choices, strings.TrimPrefix(line, "- "))
				} else if strings.HasPrefix(line, "*") {
					choices = append(choices, strings.TrimPrefix(line, "* "))
					c.Answer = i
				}
			}

			c.Choices = append(c.Choices, choices...)
			cards = append(cards, c)
			continue
		}

		ext := filepath.Ext(path)
		deckName := strings.TrimSuffix(filepath.ToSlash(path), ext)

		c := &BasicCard{
			deck:  deckName,
			front: front,
			back:  strings.Join(back, "\n"),
		}

		cards = append(cards, c)
	}

	return cards, nil
}

// `ParseDirectory` parses a directory of files into a slice of cards.
func ParseDirectory(path string) ([]Card, error) {
	var cards []Card

	file, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	for _, f := range file {
		if f.IsDir() {
			continue
		}
		if filepath.Ext(f.Name()) == ".md" {
			deck, err := ParseMarkdownFile(filepath.Join(path, f.Name()))
			if err != nil {
				return nil, err
			}
			cards = append(cards, deck...)
		}
	}

	return cards, nil
}

// `LoadDeckFromPath` loads a deck of cards from a given path.
func LoadDeckFromPath(path string) ([]Card, error) {
	var cards []Card

	ext := filepath.Ext(path)

	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("invalid path: %w", err)
	}

	if ext == ".md" {
		deck, err := ParseMarkdownFile(path)
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
