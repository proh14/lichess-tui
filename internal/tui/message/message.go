package message

import (
	"lichess-tui/internal/requests"
)

type StartGame struct {
	Time      uint
	Increment uint
}

type LoadBoard struct {
	Time      uint
	Increment uint
	Data      requests.IncomingEvents
}
