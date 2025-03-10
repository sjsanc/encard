package encard

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadCards(t *testing.T) {
	testdataDir, err := filepath.Abs("testdata")
	if err != nil {
		t.Fatalf("failed to resolve testdata directory: %v", err)
	}

	tests := []struct {
		name   string
		paths  []string
		count  int
		errors []error
	}{
		{
			name:   "empty path",
			paths:  []string{""},
			errors: []error{ErrInvalidPath},
		},
		{
			name:   "invalid path",
			paths:  []string{"invalid"},
			errors: []error{ErrInvalidPath},
		},
		{
			name:   "valid absolute path that doesn't exist",
			paths:  []string{filepath.Join(testdataDir, "loader/missing.md")},
			errors: []error{ErrInvalidPath},
		},
		{
			name:  "valid absolute path with 3 cards",
			paths: []string{filepath.Join(testdataDir, "loader/valid.md")},
			count: 3,
		},
		{
			name:  "valid absolute path with 2 valid cards and 1 invalid card",
			paths: []string{filepath.Join(testdataDir, "loader/partially_valid.md")},
			count: 2,
		},
		{
			name:  "valid absolute path to a directory loaded recursively",
			paths: []string{filepath.Join(testdataDir, "loader")},
			count: 11,
		},
		{
			name:  "valid relative path with 3 cards",
			paths: []string{"loader/valid.md"},
			count: 3,
		},
		{
			name:  "valid relative path with 2 valid cards and 1 invalid card",
			paths: []string{"loader/partially_valid.md"},
			count: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards, errors := LoadCards(tt.paths, testdataDir)

			for i, err := range errors {
				fmt.Println(i, err)
			}

			if len(tt.errors) != len(errors) {
				t.Errorf("expected %d errors, got %d", len(tt.errors), len(errors))
			}

			if tt.count != len(cards) {
				t.Errorf("expected %d cards, got %d", tt.count, len(cards))
			}
		})
	}
}

var testCard = `
# Question 1
Answer 1
---
# Question 2
Answer 2
`

func BenchmarkLoadCards(b *testing.B) {
	tmp := b.TempDir()

	for i := 0; i < 100; i++ {
		_ = os.WriteFile(tmp+"/file"+fmt.Sprintf("%d", i)+".md", []byte(testCard), 0644)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = LoadCards([]string{tmp}, "")
	}
}

// Sync
// BenchmarkLoadCards-16               1989            592441 ns/op          152429 B/op    1631 allocs/op
// PASS
// ok      github.com/sjsanc/encard/internal/encard        1.183s

// Sync + WaitGroup
// BenchmarkLoadCards-16               6831            170432 ns/op          160135 B/op    1829 allocs/op
// PASS
// ok      github.com/sjsanc/encard/internal/encard        1.169s

// SYnc + WaitGroup + Worker
// BenchmarkLoadCards-16               5335            219899 ns/op          167097 B/op       1964 allocs/op
// PASS
// ok      github.com/sjsanc/encard/internal/encard        1.205s
