package encard

import (
	"errors"
	"fmt"
	"os"

	"github.com/adrg/xdg"
	"gopkg.in/ini.v1"
)

type Config struct {
	CardsDir string
	LogsDir  string
}

var DefaultCardsDir = xdg.DataHome + "/encard/cards"
var DefaultLogsDir = xdg.StateHome + "/encard/logs"
var DefaultConfigPath = xdg.ConfigHome + "/encard/config.ini"

func NewConfig(configPath string) (*Config, error) {
	// If a config override is provided
	if configPath != "" {
		return LoadConfigFromFile(configPath)
	}
	// If a config file exists in the default location
	if _, err := os.Stat(DefaultConfigPath); err == nil {
		return LoadConfigFromFile(DefaultConfigPath)
	}
	// Otherwise, load the default configuration
	return LoadDefaultConfig()
}

// `LoadDefaultConfig` loads the default configuration.
func LoadDefaultConfig() (*Config, error) {
	err := mustExist(DefaultCardsDir)
	if err != nil {
		return nil, err
	}
	err = mustExist(DefaultLogsDir)
	if err != nil {
		return nil, err
	}

	return &Config{
		CardsDir: DefaultCardsDir,
		LogsDir:  DefaultLogsDir,
	}, nil
}

// `LoadConfigFromFile` loads a configuration file from the given path.
func LoadConfigFromFile(path string) (*Config, error) {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("config file does not exist")
	}
	if info.IsDir() {
		return nil, errors.New("config file is a directory")
	}

	cfg, err := ini.Load(path)
	if err != nil {
		return nil, err
	}

	cardsDir := DefaultCardsDir
	logsDir := DefaultLogsDir

	section, err := cfg.GetSection("storage")
	if err == nil {
		stgCards := section.Key("cards").String()
		if stgCards != "" {
			cardsDir = stgCards
		}

		stgLogs := section.Key("logs").String()
		if stgLogs != "" {
			logsDir = stgLogs
		}
	}

	err = mustExist(cardsDir)
	if err != nil {
		return nil, err
	}

	err = mustExist(logsDir)
	if err != nil {
		return nil, err
	}

	return &Config{
		CardsDir: cardsDir,
		LogsDir:  logsDir,
	}, nil
}

func mustExist(path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
	}
	return nil
}
