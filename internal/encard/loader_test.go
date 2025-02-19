package encard

import (
	"fmt"
	"os"
	"testing"

	"github.com/gobwas/glob"
)

func TestLoadCards(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		globs     []glob.Glob
		expectLen int
		expectErr bool
	}{
		{
			name:      "invalid path",
			path:      "",
			expectErr: true,
		},
		{
			name:      "valid path and no globs matches everything",
			path:      "testdata",
			expectLen: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards, err := LoadCards(tt.path, tt.globs)

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
		_, _ = LoadCards(tmp, nil)
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
