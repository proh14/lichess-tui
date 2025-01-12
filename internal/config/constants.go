package config

import (
	"os"
	"path"
)

const (
	RELATIVE_CONFIG_PATH = "lichess-tui/config.yaml"
	RELATIVE_TOKEN_PATH  = "lichesstoken"
)

// Not a constant lol
func GetConfigPath() string {
	configDir, _ := os.UserConfigDir()
	return path.Join(configDir, RELATIVE_CONFIG_PATH)
}

func AddDataDir(add string) string {
	homeDir, _ := os.UserHomeDir()
	localDir := path.Join(homeDir, ".local/share/")
	return path.Join(localDir, add)
}
