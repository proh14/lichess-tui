package config

import (
  "os"
)

const (
  REALETIVE_CONFIG_PATH = "lichess-tui/config.yaml"
)

// Not a constant lol
func GetConfigPath() string {
  configDir, _ := os.UserConfigDir()
  return configDir + REALETIVE_CONFIG_PATH
}
