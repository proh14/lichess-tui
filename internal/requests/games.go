package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"lichess-tui/internal/errors"
	"lichess-tui/internal/requests/requestTypes"
)

func SeekGame(body requestTypes.SeekGameConfig, token string) {
	bodyBytes, _ := json.Marshal(body)

	req, err := http.NewRequest(
		POST, "https://lichess.org/api/board/seek", bytes.NewBuffer(bodyBytes))
	if err != nil {
		errors.RequestError(err)
	}

	setHeaders(req, token, JSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.HandleRequestResponse(req, resp, err)
	}
}

func GetOngoingGames(token string, respVar *requestTypes.OngoingGames) {
	req, err := http.NewRequest(
		GET, "https://lichess.org/api/account/playing", nil,
	)
	if err != nil {
		errors.RequestError(err)
	}

	setHeaders(req, token, JSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.HandleRequestResponse(req, resp, err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var respStruct requestTypes.OngoingGames
	json.Unmarshal(respBody, &respStruct)

	*respVar = respStruct
}

const (
	OPERATION_RESIGN = "resign"
	OPERATION_ABORT  = "abort"
)

func GameOperation(gameId string, operation string, token string) {
	url, _ := url.JoinPath("https://lichess.org/api/board/game", gameId, operation)

	req, err := http.NewRequest(
		POST,
		url,
		nil,
	)
	if err != nil {
		errors.RequestError(err)
	}

	setHeaders(req, token, JSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.HandleRequestResponse(req, resp, err)
	}

	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)

	var respMap map[string]string
	json.Unmarshal(respBody, &respMap)

	fmt.Println(respMap)
}

func Move(gameId string, move string, body requestTypes.MoveConfig, token string) {
	url, _ := url.JoinPath("https://lichess.org/api/board/game", gameId, "move", move)
	bodyBytes, _ := json.Marshal(body)

	req, err := http.NewRequest(
		POST,
		url,
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		errors.RequestError(err)
	}

	setHeaders(req, token, JSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.HandleRequestResponse(req, resp, err)
	}

	defer resp.Body.Close()
}


var BoardStateData requestTypes.BoardState

func StreamBoardState(gameId string, token string) {
	url, _ := url.JoinPath("https://lichess.org/api/board/game/stream", gameId)

	req, err := http.NewRequest(GET, url, nil)
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
		dec.Decode(&BoardStateData)
	}
}
