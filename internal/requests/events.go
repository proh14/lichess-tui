package requests

import (
	"encoding/json"
	"net/http"

	"lichess-tui/internal/errors"
)

// https://lichess.org/api/stream/event
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
			ID         string `json:"id"`
			Username   string `json:"username"`
			Rating     uint   `json:"rating"`
			RatingDiff uint   `json:"ratingDiff"`
		} `json:"opponent"`
		Perf        string `json:"perf"`
		Rated       bool   `json:"rated"`
		SecondsLeft uint   `json:"secondsLeft"`
		Source      string `json:"source"`
		Status      struct {
			ID   uint   `json:"id"`
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
		Winner     string `json:"winner"`
		RatingDiff uint   `json:"ratingDiff"`
		ID         string `json:"id"`
	} `json:"game"`
	Challenge struct {
		ID         string `json:"id"`
		URL        string `json:"url"`
		Status     string `json:"status"`
		Challenger struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Rating uint   `json:"rating"`
			Online bool   `json:"online"`
			Lag    uint   `json:"lag"`
		} `json:"challenger"`
		DestUser struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Rating uint   `json:"rating"`
			Title  string `json:"title"`
			Online bool   `json:"online"`
			Lag    uint   `json:"lag"`
		} `json:"destUser"`
		Variant struct {
			Key   string `json:"key"`
			Name  string `json:"name"`
			Short string `json:"short"`
		} `json:"variant"`
		Rated       bool   `json:"rated"`
		Speed       string `json:"speed"`
		TimeControl struct {
			Type      string `json:"type"`
			Limit     uint   `json:"limit"`
			Increment uint   `json:"increment"`
			Show      string `json:"show"`
		} `json:"timeControl"`
		Color      string `json:"color"`
		FinalColor string `json:"finalColor"`
		Perf       struct {
			Icon string `json:"icon"`
			Name string `json:"name"`
		} `json:"perf"`
		DeclineReason    string `json:"declineReason"`
		DeclineReasonKey string `json:"declineReasonKey"`
	} `json:"challenge"`
}

var IncomingEventsData IncomingEvents

func StreamIncomingEvents(token string) {
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

	for dec.More() {
		dec.Decode(&IncomingEventsData)
	}
}
