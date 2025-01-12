package config

import (
	"os"
	"path"
)

const (
	RELATIVE_CONFIG_PATH = "lichess-tui/config.yaml"
)

// Not a constant lol
func GetConfigPath() string {
	configDir, _ := os.UserConfigDir()
	return path.Join(configDir, RELATIVE_CONFIG_PATH)
}
