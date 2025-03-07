package encard

import (
	"fmt"
	"testing"
)

func TestLoadDefaultConfig(t *testing.T) {
	cfg, err := LoadDefaultConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.CardsDir == "" {
		t.Error("expected CardsDir to be non-empty")
	}

	if cfg.LogsDir == "" {
		t.Error("expected LogsDir to be non-empty")
	}

	fmt.Println(cfg.CardsDir)
	fmt.Println(cfg.LogsDir)
}

func TestLoadConfigFromFile(t *testing.T) {
	tests := []struct {
		name           string
		cfgPath        string
		expectCardsDir string
		expectLogsDir  string
		expectErr      bool
	}{
		{
			name:           "config file does not exist",
			cfgPath:        "testdata/config/missing.ini",
			expectCardsDir: DefaultCardsDir,
			expectLogsDir:  DefaultLogsDir,
			expectErr:      true,
		},
		{
			name:           "config file is a directory",
			cfgPath:        "testdata/config",
			expectCardsDir: DefaultCardsDir,
			expectLogsDir:  DefaultLogsDir,
			expectErr:      true,
		},
		{
			name:           "custom cards directory",
			cfgPath:        "testdata/config/partial.ini",
			expectCardsDir: "/tmp/encard/cards",
			expectLogsDir:  DefaultLogsDir,
		},
		{
			name:           "custom cards and logs directories",
			cfgPath:        "testdata/config/config.ini",
			expectCardsDir: "/tmp/encard/cards",
			expectLogsDir:  "/tmp/encard/logs",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := LoadConfigFromFile(tt.cfgPath)

			if tt.expectErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if tt.expectCardsDir != cfg.CardsDir {
				t.Errorf("expected CardsDir %s, got %s", tt.expectCardsDir, cfg.CardsDir)
			}
			if tt.expectLogsDir != cfg.LogsDir {
				t.Errorf("expected LogsDir %s, got %s", tt.expectLogsDir, cfg.LogsDir)
			}
		})
	}
}
