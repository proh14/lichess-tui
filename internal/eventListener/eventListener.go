package eventListener

import (
	"lichess-tui/internal/requests"
	"lichess-tui/internal/tui/message"

	tea "github.com/charmbracelet/bubbletea"
)

func IncomingEventListener(p *tea.Program) {
	changed := false
	lastIncomingEventsData := requests.IncomingEvents{}

	for {
		if changed && requests.IncomingEventsData == lastIncomingEventsData {
			continue
		}
		switch requests.IncomingEventsData.Type {
		case "gameStart":
			msg := message.LoadBoard{
				Time:      69,
				Increment: 69,
				GameID:    requests.IncomingEventsData.Game.GameID,
			}
			p.Send(msg)
			lastIncomingEventsData = requests.IncomingEventsData
			changed = true
		}
	}
}
