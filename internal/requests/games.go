package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"github.com/proh14/lichess-tui/internal/errors"
)

// rated // bool
// time // number
// increment // number
// days // number
// variant
// ratingRange // example: 1500-1800
func SeekGame(body map[string]string, token string) {
	bodyBytes, _ := json.Marshal(body)
	req := request(POST, "https://lichess.org/api/board/seek", bytes.NewBuffer(bodyBytes))

	setHeaders(req, token, NDJSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}

	defer resp.Body.Close()
}

type OngoingGames struct {
	NowPlaying []struct {
		GameID   string `json:"gameId"`
		FullID   string `json:"fullId"`
		Color    string `json:"color"`
		Fen      string `json:"fen"`
		HasMoved bool   `json:"hasMoved"`
		IsMyTurn bool   `json:"isMyTurn"`
		LastMove string `json:"lastMove"`
		Opponent struct {
			ID       string `json:"id"`
			Rating   uint    `json:"rating"`
			Username string `json:"username"`
		} `json:"opponent"`
		Perf        string `json:"perf"`
		Rated       bool   `json:"rated"`
		SecondsLeft uint    `json:"secondsLeft"`
		Source      string `json:"source"`
		Speed       string `json:"speed"`
		Variant     struct {
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"variant"`
	} `json:"nowPlaying"`
}

func GetOngoingGames(token string) OngoingGames {
	req := request(GET, "https://lichess.org/api/account/playing", nil)

	setHeaders(req, token, NDJSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	respMap := OngoingGames{}
	json.Unmarshal(respBody, &respMap)

	return respMap
}

const (
	OPERATION_RESIGN = "resign"
	OPERATION_ABORT  = "abort"
)

func GameOperation(gameId string, operation string, token string) {
	url, _ := url.JoinPath("https://lichess.org/api/board/game", gameId, operation)

	req := request(
		POST,
		url,
		nil,
	)

	setHeaders(req, token, NDJSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}

	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)

	var respMap map[string]string
	json.Unmarshal(respBody, &respMap)

	fmt.Println(respMap)
}

