package lichess

import (
	"bytes"
	"encoding/json"
	"io"
	"path"
	"net/http"
	"github.com/proh14/lichess-tui/internal/errors"
)

const (
	GET = "GET"
	POST = "POST"
)

func setHeaders(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer " + token)
	req.Header.Set("Content-Type", "application/x-ndjson")
}

func request(method string, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		errors.RequestError(err)
	}

	return req
}

func TokenExists(token string) bool {
	req := request(GET, "https://lichess.org/api/account", nil)

	setHeaders(req, token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var respMap map[string]string
	json.Unmarshal(respBody, &respMap)

	// An error is returned in case a token doesn't exist
	_, containsKey := respMap["error"]
GetOngoingGames(token)
	return !containsKey
}

// Messages
// user string
// text string
func SendMessage(body map[string]string, token string) {
	bodyBytes, _ := json.Marshal(body)
	req := request(POST, path.Join("https://lichess.org/inbox", body["user"]), bytes.NewBuffer(bodyBytes))

	setHeaders(req, token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}

	defer resp.Body.Close()
}

// Game operations
// rated // bool
// time // number
// increment // number
// days // number
// variant 
// ratingRange // example: 1500-1800
func SeekGame(body map[string]string, token string) {
	bodyBytes, _ := json.Marshal(body)
	req := request(POST, "https://lichess.org/api/board/seek", bytes.NewBuffer(bodyBytes))

	setHeaders(req, token)

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
			Rating   int    `json:"rating"`
			Username string `json:"username"`
		} `json:"opponent"`
		Perf        string `json:"perf"`
		Rated       bool   `json:"rated"`
		SecondsLeft int    `json:"secondsLeft"`
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

	setHeaders(req, token)

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

