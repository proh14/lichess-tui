package config

import (
  "os"
)

const (
  RELATIVE_CONFIG_PATH = "lichess-tui/config.yaml"
)

// Not a constant lol
func GetConfigPath() string {
  configDir, _ := os.UserConfigDir()
  return configDir + RELATIVE_CONFIG_PATH
}
