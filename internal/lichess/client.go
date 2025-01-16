package lichess

import (
	"errors"

	"github.com/proh14/lichess-tui/internal/requests"
)

func ValidateToken(token string) error {
	if len(token) == 0 {
		return errors.New("The token can't be empty.")
	}

	if requests.TokenExists(token) {
		return nil
	}

	return errors.New("The given token is invalid.")
}
