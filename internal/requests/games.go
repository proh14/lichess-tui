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
type SeekGameConfig struct {
	Rated       bool   `json:"bool,omitempty"`
	Time        uint   `json:"time"`
	Increment   uint   `json:"increment"`
	Days        uint   `json:"days,omitempty"`
	Variant     string `json:"variant,omitempty"`
	RatingRange string `json:"ratingRange,omitempty"`
}

func SeekGame(body SeekGameConfig, token string) {
	bodyBytes, _ := json.Marshal(body)
	req := request(POST, "https://lichess.org/api/board/seek", bytes.NewBuffer(bodyBytes))

	setHeaders(req, token, JSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}

	defer resp.Body.Close()
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

// https://lichess.org/api#tag/Board/operation/boardGameStream
type BoardState struct {
	ID      string `json:"id,omitempty"`
	Variant struct {
		Key   string `json:"key,omitempty"`
		Name  string `json:"name,omitempty"`
		Short string `json:"short,omitempty"`
	} `json:"variant,omitempty"`
	Speed string `json:"speed,omitempty"`
	Perf  struct {
		Name string `json:"name,omitempty"`
	} `json:"perf,omitempty"`
	Rated     bool   `json:"rated,omitempty"`
	CreatedAt uint64 `json:"createdAt,omitempty"`
	White     struct {
		ID     string `json:"id,omitempty"`
		Name   string `json:"name,omitempty"`
		Title  string `json:"title,omitempty"`
		Rating uint   `json:"rating,omitempty"`
	} `json:"white,omitempty"`
	Black struct {
		ID          string `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		Title       string `json:"title,omitempty"`
		Rating      uint   `json:"rating,omitempty"`
		Provisional bool   `json:"provisional,omitempty"`
	} `json:"black,omitempty"`
	InitialFen string `json:"initialFen,omitempty"`
	Clock      struct {
		Initial   uint `json:"initial,omitempty"`
		Increment uint `json:"increment,omitempty"`
	} `json:"clock,omitempty"`
	Type  string `json:"type,omitempty"`
	State struct {
		Type   string `json:"type,omitempty"`
		Moves  string `json:"moves,omitempty"`
		Wtime  uint   `json:"wtime,omitempty"`
		Btime  uint   `json:"btime,omitempty"`
		Winc   uint   `json:"winc,omitempty"`
		Binc   uint   `json:"binc,omitempty"`
		Status string `json:"status,omitempty"`
	} `json:"state,omitempty"`
	Moves             string `json:"moves,omitempty"`
	Wtime             uint   `json:"wtime,omitempty"`
	Btime             uint   `json:"btime,omitempty"`
	Winc              uint   `json:"winc,omitempty"`
	Binc              uint   `json:"binc,omitempty"`
	Wdraw             bool   `json:"wdraw,omitempty"`
	Bdraw             bool   `json:"bdraw,omitempty"`
	Status            string `json:"status,omitempty"`
	Username          string `json:"username,omitempty"`
	Text              string `json:"text,omitempty"`
	Room              string `json:"room,omitempty"`
	Gone              bool   `json:"gone,omitempty"`
	ClaimWinInSeconds uint   `json:"claimWinInSeconds,omitempty"`
}

var BoardStateData BoardState

func StreamBoardState(gameId string, token string) {
	url, _ := url.JoinPath("https://lichess.org/api/board/game/stream", gameId)
	req := request(GET, url, nil)

	setHeaders(req, token, NDJSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	for dec.More() {
		dec.Decode(&BoardStateData)
	}
}
