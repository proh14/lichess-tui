package lichess

import (
	"errors"
)

func ValidateToken(value string) error {
	if len(value) == 0 {
		return errors.New("The token can't be empty.")
	}

	if TokenExists(value) {
		return nil
	}

	return errors.New("The given token is invalid.")
}

