package encard

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadCards(t *testing.T) {
	tests := []struct {
		name   string
		paths  []string
		cfg    string
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
			paths:  []string{"testdata/loader/invalid"},
			errors: []error{ErrInvalidPath},
		},
		{
			name:  "valid absolute path with 3 cards",
			paths: []string{"testdata/loader/one.md"},
			count: 3,
		},
		{
			name:  "valid absolute path with 2 valid cards and 1 invalid card",
			paths: []string{"testdata/loader/partial.md"},
			count: 2,
		},
		{
			name:  "valid relative path with 3 cards",
			paths: []string{"one.md"},
			cfg:   "testdata/loader/config.ini",
			count: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := NewConfig(tt.cfg)
			if err != nil {
				t.Fatalf("error loading config: %v", err)
			}

			cards, errors := LoadCards(tt.paths, cfg)

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
		cfg, _ := NewConfig("")
		_, _ = LoadCards([]string{tmp}, cfg)
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
