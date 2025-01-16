package lichess

import (
	"errors"

	"github.com/proh14/lichess-tui/internal/requests"
)

func ValidateToken(value string) error {
	if len(value) == 0 {
		return errors.New("The token can't be empty.")
	}

	if requests.TokenExists(value) {
		return nil
	}

	return errors.New("The given token is invalid.")
}
