package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"lichess-tui/internal/errors"
)

// https://lichess.org/api#tag/Board/operation/apiBoardSeek
type SeekGameResponse struct {
	Id    string `json:"id"`
}

// https://lichess.org/api#tag/Board/operation/apiBoardSeek
type SeekGameConfig struct {
	Rated       bool   `json:"bool,omitempty"`
	Time        uint   `json:"time"`
	Increment   uint   `json:"increment,omitempty"`
	Days        uint   `json:"days,omitempty"`
	Variant     string `json:"variant,omitempty"`
	RatingRange string `json:"ratingRange,omitempty"`
}

func SeekGame(body SeekGameConfig, token string, respVar *SeekGameResponse) {
	bodyBytes, _ := json.Marshal(body)
	req := request(POST, "https://lichess.org/api/board/seek", bytes.NewBuffer(bodyBytes))

	setHeaders(req, token, NDJSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}

	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var respStruct SeekGameResponse
	json.Unmarshal(respBody, &respStruct)

	*respVar = respStruct
}

// https://lichess.org/api#tag/Account/operation/accountMe
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
			Rating   uint   `json:"rating"`
			Username string `json:"username"`
		} `json:"opponent"`
		Perf        string `json:"perf"`
		Rated       bool   `json:"rated"`
		SecondsLeft uint   `json:"secondsLeft"`
		Source      string `json:"source"`
		Speed       string `json:"speed"`
		Variant     struct {
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"variant"`
	} `json:"nowPlaying"`
}

func GetOngoingGames(token string, respVar *OngoingGames) {
	req := request(GET, "https://lichess.org/api/account/playing", nil)

	setHeaders(req, token, NDJSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var respStruct OngoingGames
	json.Unmarshal(respBody, &respStruct)

	*respVar = respStruct
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

// https://lichess.org/api#tag/Board/operation/boardGameMove
type MoveConfig struct {
	OfferingDraw uint `json:"offeringDraw,omitempty"`
}

func Move(gameId string, move string, body MoveConfig, token string) {
	url, _ := url.JoinPath("https://lichess.org/api/board/game", gameId, "move", move)
	bodyBytes, _ := json.Marshal(body)

	req := request(
		POST,
		url,
		bytes.NewBuffer(bodyBytes),
	)

	setHeaders(req, token, NDJSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}

	defer resp.Body.Close()
}
