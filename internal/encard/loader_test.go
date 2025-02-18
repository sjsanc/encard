package encard

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestLoadDeckFromPath(t *testing.T) {
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stderr)

	tests := []struct {
		name      string
		path      string
		setupFn   func()
		expectErr bool
		expectLen int
	}{
		{
			name:      "invalid path",
			path:      "",
			expectErr: true,
		},
		{
			name:      "valid path and invalid extension",
			path:      "testdata/invalid.txt",
			expectErr: true,
		},
		{
			name:      "valid path and empty markdown file",
			path:      "testdata/md/empty.md",
			expectErr: true,
		},
		{
			name:      "valid path and empty directory",
			path:      "testdata/empty",
			expectErr: true,
		},
		{
			name:      "valid path and partially valid markdown file",
			path:      "testdata/md/partial.md",
			expectLen: 2,
		},
		{
			name:      "valid path and valid markdown file",
			path:      "testdata/md/valid.md",
			expectLen: 3,
		},
		{
			name:      "valid path and valid markdown directory",
			path:      "testdata/md/dir",
			expectLen: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards, err := LoadDeckFromPath(tt.path)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			len := len(cards)

			if tt.expectLen != len {
				t.Errorf("expected %d cards, got %d", tt.expectLen, len)
			}
		})
	}
}
