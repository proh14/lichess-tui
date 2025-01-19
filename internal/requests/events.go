package requests

import (
	"encoding/json"
	"net/http"
	"lichess-tui/internal/errors"
)

type IncomingEvents struct {
	Type string `json:"type"`
	Game struct {
		GameID   string `json:"gameId"`
		FullID   string `json:"fullId"`
		Color    string `json:"color"`
		Fen      string `json:"fen"`
		HasMoved bool   `json:"hasMoved"`
		IsMyTurn bool   `json:"isMyTurn"`
		LastMove string `json:"lastMove"`
		Opponent struct {
			ID       string `json:"id"`
			Rating   int    `json:"rating"`
			Username string `json:"username"`
		} `json:"opponent"`
		Perf        string `json:"perf"`
		Rated       bool   `json:"rated"`
		SecondsLeft int    `json:"secondsLeft"`
		Source      string `json:"source"`
		Status      struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"status"`
		Speed   string `json:"speed"`
		Variant struct {
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"variant"`
		Compat struct {
			Bot   bool `json:"bot"`
			Board bool `json:"board"`
		} `json:"compat"`
		ID string `json:"id"`
	} `json:"game"`
}

var IncomingEventsData IncomingEvents

func StreamIncomingEvents(token string) {
	req := request(GET, "https://lichess.org/api/stream/event", nil)

	setHeaders(req, token, NDJSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	
	for {
		dec.Decode(&IncomingEventsData)
	}
}

