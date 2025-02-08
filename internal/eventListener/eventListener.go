// package eventListener
//
// import (
// 	"lichess-tui/internal/requests"
// 	"lichess-tui/internal/tui/message"
//
// 	tea "github.com/charmbracelet/bubbletea"
// )

// func IncomingEventListener(p *tea.Program) {
// 	changed := false
// 	lastIncomingEventsData := requests.IncomingEvents{}
//
// 	for {
// 		if changed && requests.IncomingEventsData == lastIncomingEventsData {
// 			continue
// 		}
// 		switch {
// 		case requests.IncomingEventsData.Type == "gameStart" && requests.IncomingEventsData.Game.Fen != "" && requests.IncomingEventsData.Game.GameID != "":
// 			msg := message.LoadBoard{
// 				Time:      69,
// 				Increment: 69,
// 				Data:      requests.IncomingEventsData,
// 			}
// 			p.Send(msg)
// 			lastIncomingEventsData = requests.IncomingEventsData
// 			changed = true
// 		}
// 	}
// }
