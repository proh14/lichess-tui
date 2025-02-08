package message

import (
	"lichess-tui/internal/requests/requestTypes"
)

type StartGame struct {
	Time      uint
	Increment uint
}

type LoadBoard struct {
	Time      uint
	Increment uint
	Data      requestTypes.IncomingEvents
}
