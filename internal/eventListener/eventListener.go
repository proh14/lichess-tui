package eventListener

import (
	"lichess-tui/internal/requests"
	"lichess-tui/internal/tui/message"

	tea "github.com/charmbracelet/bubbletea"
)

func IncomingEventListener(p *tea.Program) {
	// lastIncomingEventsData := requests.IncomingEvents{}
	for {
		// if requests.IncomingEventsData == lastIncomingEventsData {
		// 	continue
		// }
		switch requests.IncomingEventsData.Type {
		case "gameStart":
			msg := message.LoadBoard{Time: 69, Increment: 69}
			p.Send(msg)
			// lastIncomingEventsData = requests.IncomingEventsData
		}
	}
}
