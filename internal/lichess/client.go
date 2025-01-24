package lichess

import (
	"errors"

	"lichess-tui/internal/requests"
)

func ValidateToken(token string, configPath string) error {
	if len(token) == 0 {
		return errors.New("The token can't be empty.")
	}

	if requests.TokenExists(token) {
		return nil
	}

	return errors.New(
		"The given token is invalid. " + "Please try to remove " + configPath + " and log in again.",
	)
}
