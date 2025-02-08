package requests

import (
	"encoding/json"
	"net/http"

	"lichess-tui/internal/errors"
	"lichess-tui/internal/tui/message"
	"lichess-tui/internal/requests/requestTypes"

	tea "github.com/charmbracelet/bubbletea"
)

var IncomingEventsData requestTypes.IncomingEvents

func StreamIncomingEvents(token string, p *tea.Program) {
	req, err := http.NewRequest(
		GET, "https://lichess.org/api/stream/event", nil)
	if err != nil {
		errors.RequestError(err)
	}

	setHeaders(req, token, NDJSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.HandleRequestResponse(req, resp, err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	lastIncomingEventsData := requestTypes.IncomingEvents{}
	changed := false

	for dec.More() {
		dec.Decode(&IncomingEventsData)
		if !changed || IncomingEventsData != lastIncomingEventsData {
			lastIncomingEventsData = IncomingEventsData
			changed = true

			switch {
			case IncomingEventsData.Type == "gameStart":
				msg := message.LoadBoard{
					Time:      69,
					Increment: 69,
					Data:      IncomingEventsData,
				}
				p.Send(msg)
			}
		}
	}
}
