package encard

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestResolveRootPath(t *testing.T) {
	tmp := t.TempDir()

	tests := []struct {
		name           string
		envXdgDataHome string
		envHome        string
		expectErr      bool
		expectPath     string
	}{
		{
			name:           "neither XDG_DATA_HOME nor HOME are set",
			envXdgDataHome: "",
			envHome:        "",
			expectErr:      true,
		},
		{
			name:           "XDG_DATA_HOME is set",
			envXdgDataHome: tmp + "/xdg_data_home",
			envHome:        "",
			expectPath:     tmp + "/xdg_data_home/encard",
		},
		{
			name:           "XDG_DATA_HOME is not set and HOME is set",
			envXdgDataHome: "",
			envHome:        tmp + "/home",
			expectPath:     tmp + "/home/encard",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("XDG_DATA_HOME", tt.envXdgDataHome)
			t.Setenv("HOME", tt.envHome)

			path, err := ResolveRootPath()

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			if tt.expectPath != path {
				t.Errorf("expected path %s, got %s", tt.expectPath, path)
			}

			info, err := os.Stat(path)
			if err != nil {
				t.Errorf("expected path %s to exist, got error: %v", path, err)
			}
			if !info.IsDir() {
				t.Errorf("expected path %s to be a directory, got a file", path)
			}

			t.Cleanup(func() {
				os.RemoveAll(path)
			})
		})
	}
}

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
