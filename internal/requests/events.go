package requests

import (
	"encoding/json"
	"net/http"

	"lichess-tui/internal/errors"
)

// https://lichess.org/api/stream/event
type IncomingEvents struct {
	Type string `json:"type,omitempty"`
	Game struct {
		GameID   string `json:"gameId,omitempty"`
		FullID   string `json:"fullId,omitempty"`
		Color    string `json:"color,omitempty"`
		Fen      string `json:"fen,omitempty"`
		HasMoved bool   `json:"hasMoved,omitempty"`
		IsMyTurn bool   `json:"isMyTurn,omitempty"`
		LastMove string `json:"lastMove,omitempty"`
		Opponent struct {
			ID         string `json:"id,omitempty"`
			Username   string `json:"username,omitempty"`
			Rating     uint   `json:"rating,omitempty"`
			RatingDiff uint   `json:"ratingDiff,omitempty"`
		} `json:"opponent,omitempty"`
		Perf        string `json:"perf,omitempty"`
		Rated       bool   `json:"rated,omitempty"`
		SecondsLeft uint   `json:"secondsLeft,omitempty"`
		Source      string `json:"source,omitempty"`
		Status      struct {
			ID   uint   `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"status,omitempty"`
		Speed   string `json:"speed,omitempty"`
		Variant struct {
			Key  string `json:"key,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"variant,omitempty"`
		Compat struct {
			Bot   bool `json:"bot,omitempty"`
			Board bool `json:"board,omitempty"`
		} `json:"compat,omitempty"`
		Winner     string `json:"winner,omitempty"`
		RatingDiff uint   `json:"ratingDiff,omitempty"`
		ID         string `json:"id,omitempty"`
	} `json:"game,omitempty"`
	Challenge struct {
		ID         string `json:"id,omitempty"`
		URL        string `json:"url,omitempty"`
		Status     string `json:"status,omitempty"`
		Challenger struct {
			ID     string `json:"id,omitempty"`
			Name   string `json:"name,omitempty"`
			Rating uint   `json:"rating,omitempty"`
			Online bool   `json:"online,omitempty"`
			Lag    uint   `json:"lag,omitempty"`
		} `json:"challenger,omitempty"`
		DestUser struct {
			ID     string `json:"id,omitempty"`
			Name   string `json:"name,omitempty"`
			Rating uint   `json:"rating,omitempty"`
			Title  string `json:"title,omitempty"`
			Online bool   `json:"online,omitempty"`
			Lag    uint   `json:"lag,omitempty"`
		} `json:"destUser,omitempty"`
		Variant struct {
			Key   string `json:"key,omitempty"`
			Name  string `json:"name,omitempty"`
			Short string `json:"short,omitempty"`
		} `json:"variant,omitempty"`
		Rated       bool   `json:"rated,omitempty"`
		Speed       string `json:"speed,omitempty"`
		TimeControl struct {
			Type      string `json:"type,omitempty"`
			Limit     uint   `json:"limit,omitempty"`
			Increment uint   `json:"increment,omitempty"`
			Show      string `json:"show,omitempty"`
		} `json:"timeControl,omitempty"`
		Color      string `json:"color,omitempty"`
		FinalColor string `json:"finalColor,omitempty"`
		Perf       struct {
			Icon string `json:"icon,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"perf,omitempty"`
		DeclineReason    string `json:"declineReason,omitempty"`
		DeclineReasonKey string `json:"declineReasonKey,omitempty"`
	} `json:"challenge,omitempty"`
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
