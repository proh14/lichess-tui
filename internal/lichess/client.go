package lichess

import (
	"errors"
	"os"

	"lichess-tui/internal/requests"
)

func ValidateToken(token string, configPath string) error {
	if len(token) == 0 {
		return errors.New("The token can't be empty.")
	}

	if requests.TokenExists(token) {
		return nil
	}

	message := "The given token is invalid. "

	_, err := os.Stat(configPath)
	if !os.IsNotExist(err) {
		message += "Please try to remove " + configPath + " and log in again."
	}

	return errors.New(
		message,
	)
}
